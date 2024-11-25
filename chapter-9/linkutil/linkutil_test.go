package linkutil_test

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"programming-in-Go-by-Mark-Summerfield/chapter-9/linkutil"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestLinksFromReader(t *testing.T) {
	html, err := os.Open("./index.html")
	if err != nil {
		t.Fatal(err)
	}

	defer html.Close()

	var links []string
	if links, err = linkutil.LinksFromReader(html); err != nil {
		t.Fatal(err)
	}

	expectedFile, err := os.ReadFile("./expected.txt")
	buffer := bufio.NewReader(bytes.NewBuffer(expectedFile))

	var expected []string
	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}

			break
		}

		line = strings.TrimSpace(line)
		expected = append(expected, line)
	}

	sort.Strings(expected)
	sort.Strings(links)
	if !reflect.DeepEqual(links, expected) {
		t.Errorf("expected:\n%s\n\ngot:\n%s", strings.Join(expected, "\n"), strings.Join(links, "\n"))
	}
}
