package Parser

import "encoding/xml"

/*
LibraryDirectory is struct containing information about plex library
*/
type LibraryDirectory struct {
	Key  int    `xml:"key,attr"`
	Lang string `xml:"language,attr"`
}

/*
ParseLibraries parses libraries XML returned by plex
See: https://github.com/mjs7231/python-plexapi/wiki/Unofficial-Plex-API-Documentation
*/
func ParseLibraries(librariesXML []byte) []LibraryDirectory {
	type libraryParserQuery struct {
		Directories []LibraryDirectory `xml:"Directory"`
	}
	var q libraryParserQuery
	xml.Unmarshal(librariesXML, &q)

	return q.Directories
}
