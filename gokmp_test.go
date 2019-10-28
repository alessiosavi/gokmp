package gokmp

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"strings"
	"testing"
)

// TESTS

// pretty much the worst case string for strings.Index
// wrt. string and pattern
const str = "aabaabaaaabbaabaabaaabbaabaabb"
const pattern = "aabb"
const dantePattern = "stelle"

func readTextData(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	data := string(content)
	return data
}

func TestFindAllStringIndex(t *testing.T) {
	kmp, _ := NewKMP(pattern)
	// fmt.Println(kmp)
	ints := kmp.FindAllStringIndex(str)
	test := []int{8, 19, 26}
	for i, v := range ints {
		if test[i] != v {
			t.Errorf("FindAllStringIndex:\t%v != %v, (exp: %v != act: %v)", test[i], v, ints, test)
		}
	}
}

func TestFindStringIndex(t *testing.T) {
	kmp, _ := NewKMP(pattern)
	ints := kmp.FindStringIndex(str)
	test := 8
	if ints != test {
		t.Errorf("FindStringIndex:\t%v != %v", ints, test)
	}
}

func TestContainedIn(t *testing.T) {
	kmp, _ := NewKMP(pattern)
	if !kmp.ContainedIn(str) {
		t.Errorf("ContainedIn:\tExpected: True !=  actual: False")
	}
}

func TestOccurrences(t *testing.T) {
	kmp, _ := NewKMP(pattern)
	nr := kmp.Occurrences(str)
	if nr != 3 {
		t.Errorf("Occurences:\texp: %v != act: %v)", 3, nr)
	}
}

func TestOccurrencesFail(t *testing.T) {
	kmp, _ := NewKMP(pattern)
	nr := kmp.Occurrences("pebble")
	if nr != 0 {
		t.Errorf("Occurences:\texp: %v != act: %v)", 0, nr)
	}
}

// BENCHMARKS

func BenchmarkKMPIndexComparison(b *testing.B) {
	kmp, _ := NewKMP(pattern)
	for i := 0; i < b.N; i++ {
		_ = kmp.FindStringIndex(str)
	}
}

func BenchmarkStringsIndexComparison(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Index(str, pattern)
	}
}

func BenchmarkKMPIndexComparisonDanteKO(b *testing.B) {
	data := readTextData("/tmp/dante.txt")
	b.ResetTimer()
	kmp, _ := NewKMP(pattern)
	for i := 0; i < b.N; i++ {
		_ = kmp.FindStringIndex(data)
	}
}

func BenchmarkStringsIndexComparisonDanteKO(b *testing.B) {
	data := readTextData("/tmp/dante.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Index(data, pattern)
	}
}

func BenchmarkKMPIndexComparisonDanteOK(b *testing.B) {
	data, err := ReadZip("dante.zip")
	if err != nil {
		b.Fail()
	}
	text := data["dante.txt"]
	b.ResetTimer()
	kmp, _ := NewKMP(dantePattern)
	for i := 0; i < b.N; i++ {
		_ = kmp.FindStringIndex(text)
	}
}

func BenchmarkStringsIndexComparisonDanteOK(b *testing.B) {
	data, err := ReadZip("dante.zip")
	if err != nil {
		b.Fail()
	}
	text := data["dante.txt"]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Index(text, dantePattern)
	}
}

func Test(t *testing.T) {
	data, err := ReadZip("dante.zip")
	if err != nil {
		t.Error(err)
	}
	for item := range data {
		t.Log("Key", item)
	}
}

// ReadZip is delegated to extract the files and read the content
func ReadZip(filename string) (map[string]string, error) {
	var filesContent map[string]string

	//log.Println("ReadZip | Opening zipped content from [" + filename + "]")
	zf, err := zip.OpenReader(filename)
	if err != nil {
		//log.Println("ReadZip | Error during read "+filename+" | Err: ", err)
		return nil, err
	}
	defer zf.Close()
	filesContent = make(map[string]string)
	for _, file := range zf.File {
		if file.Mode().IsRegular() {
			//log.Println("ReadZip | Unzipping regular file " + file.Name)
			data, err := ReadZipFile(file)
			if err != nil {
				//log.Println("ReadZip | Unable to unzip file " + file.Name)
			} else {
				//log.Println("ReadZip | File unzipped successfully!")
				filesContent[file.Name] = data
			}
		}
	}
	//log.Println("ReadZip | Unzipped ", len(filesContent), " files")
	return filesContent, nil
}

// ReadZipFile is a wrapper function for ioutil.ReadAll. It accepts a zip.File as
// its parameter, opens it, reads its content and returns it as a byte slice.
func ReadZipFile(file *zip.File) (string, error) {
	if !file.Mode().IsRegular() {
		return "", errors.New("ReadZipFile | File " + file.Name + " is not a regular!")
	}
	fc, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fc.Close()

	content, err := ioutil.ReadAll(fc)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
