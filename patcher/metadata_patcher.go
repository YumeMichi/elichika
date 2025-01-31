// Go implementation of https://github.com/JeremieCHN/MetaDataStringEditor/blob/master/MetadataFile.cs
package patcher

import (
	"elichika/config"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

type MetadataFile struct {
	reader                  *os.File
	stringLiteralOffset     uint32
	stringLiteralCount      uint32
	dataInfoPosition        int64
	stringLiteralDataOffset uint32
	stringLiteralDataCount  uint32
	stringLiterals          []StringLiteral
	strBytes                [][]byte
}

type StringLiteral struct {
	Length uint32
	Offset uint32
}

// NewMetadataFile opens the file and initializes the MetadataFile structure
func NewMetadataFile(filename string) (*MetadataFile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	meta := &MetadataFile{
		reader:         file,
		stringLiterals: make([]StringLiteral, 0),
		strBytes:       make([][]byte, 0),
	}

	// Read the file header
	if err := meta.ReadHeader(); err != nil {
		return nil, err
	}

	// Read string literals
	if err := meta.ReadLiteral(); err != nil {
		return nil, err
	}

	// Read the string bytes
	if err := meta.ReadStrByte(); err != nil {
		return nil, err
	}

	fmt.Println("Read metadata file complete")
	return meta, nil
}

// ReadHeader reads the metadata file header
func (meta *MetadataFile) ReadHeader() error {
	var signature uint32
	err := binary.Read(meta.reader, binary.LittleEndian, &signature)
	if err != nil {
		return err
	}

	if signature != 0xFAB11BAF {
		return fmt.Errorf("signature check failed")
	}

	var version int32
	binary.Read(meta.reader, binary.LittleEndian, &version)
	binary.Read(meta.reader, binary.LittleEndian, &meta.stringLiteralOffset)
	binary.Read(meta.reader, binary.LittleEndian, &meta.stringLiteralCount)
	meta.dataInfoPosition, _ = meta.reader.Seek(0, io.SeekCurrent)
	binary.Read(meta.reader, binary.LittleEndian, &meta.stringLiteralDataOffset)
	binary.Read(meta.reader, binary.LittleEndian, &meta.stringLiteralDataCount)

	return nil
}

// ReadLiteral reads the string literal list
func (meta *MetadataFile) ReadLiteral() error {
	_, err := meta.reader.Seek(int64(meta.stringLiteralOffset), io.SeekStart)
	if err != nil {
		return err
	}

	for i := 0; i < int(meta.stringLiteralCount)/8; i++ {
		var lit StringLiteral
		binary.Read(meta.reader, binary.LittleEndian, &lit.Length)
		binary.Read(meta.reader, binary.LittleEndian, &lit.Offset)
		meta.stringLiterals = append(meta.stringLiterals, lit)
	}

	return nil
}

// ReadStrByte reads the string byte arrays
func (meta *MetadataFile) ReadStrByte() error {
	for _, literal := range meta.stringLiterals {
		_, err := meta.reader.Seek(int64(meta.stringLiteralDataOffset+literal.Offset), io.SeekStart)
		if err != nil {
			return err
		}

		bytes := make([]byte, literal.Length)
		_, err = meta.reader.Read(bytes)
		if err != nil {
			return err
		}
		meta.strBytes = append(meta.strBytes, bytes)
	}

	return nil
}

// FindAndReplace finds the target string and replaces it with a new string
func (meta *MetadataFile) FindAndReplace(target string, replacement string) {
	for i, bytes := range meta.strBytes {
		str := string(bytes)
		if strings.Contains(str, target) {
			// Print the address of the found string (its offset)
			fmt.Printf("Found string %s at offset: 0x%X\n", str, meta.stringLiterals[i].Offset)

			// Perform replacement
			str = strings.Replace(str, target, replacement, -1)
			meta.strBytes[i] = []byte(str) // Update the byte array with the new string

			// Print the new length after replacement
			fmt.Printf("Replaced string to %s at offset: 0x%X\n", replacement, meta.stringLiterals[i].Offset)
		}
	}
}

// WriteToNewFile writes the updated data to a new file
func (meta *MetadataFile) WriteToNewFile(filename string) error {
	writer, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer writer.Close()

	// Copy the entire content of the original file
	meta.reader.Seek(0, io.SeekStart)
	if _, err := io.Copy(writer, meta.reader); err != nil {
		return err
	}

	// Update string literals in the new file
	writer.Seek(int64(meta.stringLiteralOffset), io.SeekStart)
	var count uint32 = 0
	for i, literal := range meta.stringLiterals {
		literal.Offset = count
		literal.Length = uint32(len(meta.strBytes[i]))
		binary.Write(writer, binary.LittleEndian, literal.Length)
		binary.Write(writer, binary.LittleEndian, literal.Offset)
		count += literal.Length
	}

	// Align to 4 bytes if necessary
	if padding := (meta.stringLiteralDataOffset + count) % 4; padding != 0 {
		count += 4 - padding
	}

	// Ensure there is enough space for string data
	if count > meta.stringLiteralDataCount {
		fileInfo, err := writer.Stat()
		if err != nil {
			return err
		}

		if meta.stringLiteralDataOffset+meta.stringLiteralDataCount < uint32(fileInfo.Size()) {
			meta.stringLiteralDataOffset = uint32(fileInfo.Size())
		}
	}
	meta.stringLiteralDataCount = count

	// Write the string bytes
	writer.Seek(int64(meta.stringLiteralDataOffset), io.SeekStart)
	for _, str := range meta.strBytes {
		writer.Write(str)
	}

	// Update the header with the new offsets and lengths
	writer.Seek(meta.dataInfoPosition, io.SeekStart)
	binary.Write(writer, binary.LittleEndian, meta.stringLiteralDataOffset)
	binary.Write(writer, binary.LittleEndian, meta.stringLiteralDataCount)

	fmt.Println("Update metadata file complete")
	return nil
}

func MetadataPatcher(inFile, outFile string, patcher []config.Patcher) error {
	// Replace "global-metadata.dat" with your actual file name
	meta, err := NewMetadataFile(inFile)
	if err != nil {
		return err
	}
	defer meta.reader.Close()

	// Replace server URL
	// meta.FindAndReplace("https://api.garupa.jp/api/", "http://192.168.1.123:8080/")
	for _, patch := range patcher {
		meta.FindAndReplace(patch.Target, patch.Replacement)
	}

	// Write to new file
	if err := meta.WriteToNewFile(outFile); err != nil {
		return err
	}

	return nil
}
