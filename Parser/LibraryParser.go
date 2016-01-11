package Parser

import "encoding/xml"

/*
LibraryItem is struct containing information about plex library
*/
type LibraryItem struct {
	Key  int    `xml:"key,attr"`
	Lang string `xml:"language,attr"`
}

/*
ParseItems parses library items XML returned by plex
See: https://github.com/mjs7231/python-plexapi/wiki/Unofficial-Plex-API-Documentation
*/
func ParseItems(itemsXML []byte) []LibraryItem {
	type libraryParserQuery struct {
		Directories []LibraryItem `xml:"Directory"`
	}
	var q libraryParserQuery
	xml.Unmarshal(itemsXML, &q)

	return q.Directories
}
