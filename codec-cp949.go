
package charmap

type codecCP949 struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecCP949) Encode(s string) string {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecCP949) Decode(s string) string {
	return mapBytesToRunes(c.DecodeMap, s)
}

func init() {
	charmapDecode := map[byte]rune{
		'\x00':	'\u0000',	 // NULL
		'\x01':	'\u0001',	 // START OF HEADING
		'\x02':	'\u0002',	 // START OF TEXT
		'\x03':	'\u0003',	 // END OF TEXT
		'\x04':	'\u0004',	 // END OF TRANSMISSION
		'\x05':	'\u0005',	 // ENQUIRY
		'\x06':	'\u0006',	 // ACKNOWLEDGE
		'\x07':	'\u0007',	 // BELL
		'\x08':	'\u0008',	 // BACKSPACE
		'\x09':	'\u0009',	 // HORIZONTAL TABULATION
		'\x0A':	'\u000A',	 // LINE FEED
		'\x0B':	'\u000B',	 // VERTICAL TABULATION
		'\x0C':	'\u000C',	 // FORM FEED
		'\x0D':	'\u000D',	 // CARRIAGE RETURN
		'\x0E':	'\u000E',	 // SHIFT OUT
		'\x0F':	'\u000F',	 // SHIFT IN
		'\x10':	'\u0010',	 // DATA LINK ESCAPE
		'\x11':	'\u0011',	 // DEVICE CONTROL ONE
		'\x12':	'\u0012',	 // DEVICE CONTROL TWO
		'\x13':	'\u0013',	 // DEVICE CONTROL THREE
		'\x14':	'\u0014',	 // DEVICE CONTROL FOUR
		'\x15':	'\u0015',	 // NEGATIVE ACKNOWLEDGE
		'\x16':	'\u0016',	 // SYNCHRONOUS IDLE
		'\x17':	'\u0017',	 // END OF TRANSMISSION BLOCK
		'\x18':	'\u0018',	 // CANCEL
		'\x19':	'\u0019',	 // END OF MEDIUM
		'\x1A':	'\u001A',	 // SUBSTITUTE
		'\x1B':	'\u001B',	 // ESCAPE
		'\x1C':	'\u001C',	 // FILE SEPARATOR
		'\x1D':	'\u001D',	 // GROUP SEPARATOR
		'\x1E':	'\u001E',	 // RECORD SEPARATOR
		'\x1F':	'\u001F',	 // UNIT SEPARATOR
		'\x20':	'\u0020',	 // SPACE
		'\x21':	'\u0021',	 // EXCLAMATION MARK
		'\x22':	'\u0022',	 // QUOTATION MARK
		'\x23':	'\u0023',	 // NUMBER SIGN
		'\x24':	'\u0024',	 // DOLLAR SIGN
		'\x25':	'\u0025',	 // PERCENT SIGN
		'\x26':	'\u0026',	 // AMPERSAND
		'\x27':	'\u0027',	 // APOSTROPHE
		'\x28':	'\u0028',	 // LEFT PARENTHESIS
		'\x29':	'\u0029',	 // RIGHT PARENTHESIS
		'\x2A':	'\u002A',	 // ASTERISK
		'\x2B':	'\u002B',	 // PLUS SIGN
		'\x2C':	'\u002C',	 // COMMA
		'\x2D':	'\u002D',	 // HYPHEN-MINUS
		'\x2E':	'\u002E',	 // FULL STOP
		'\x2F':	'\u002F',	 // SOLIDUS
		'\x30':	'\u0030',	 // DIGIT ZERO
		'\x31':	'\u0031',	 // DIGIT ONE
		'\x32':	'\u0032',	 // DIGIT TWO
		'\x33':	'\u0033',	 // DIGIT THREE
		'\x34':	'\u0034',	 // DIGIT FOUR
		'\x35':	'\u0035',	 // DIGIT FIVE
		'\x36':	'\u0036',	 // DIGIT SIX
		'\x37':	'\u0037',	 // DIGIT SEVEN
		'\x38':	'\u0038',	 // DIGIT EIGHT
		'\x39':	'\u0039',	 // DIGIT NINE
		'\x3A':	'\u003A',	 // COLON
		'\x3B':	'\u003B',	 // SEMICOLON
		'\x3C':	'\u003C',	 // LESS-THAN SIGN
		'\x3D':	'\u003D',	 // EQUALS SIGN
		'\x3E':	'\u003E',	 // GREATER-THAN SIGN
		'\x3F':	'\u003F',	 // QUESTION MARK
		'\x40':	'\u0040',	 // COMMERCIAL AT
		'\x41':	'\u0041',	 // LATIN CAPITAL LETTER A
		'\x42':	'\u0042',	 // LATIN CAPITAL LETTER B
		'\x43':	'\u0043',	 // LATIN CAPITAL LETTER C
		'\x44':	'\u0044',	 // LATIN CAPITAL LETTER D
		'\x45':	'\u0045',	 // LATIN CAPITAL LETTER E
		'\x46':	'\u0046',	 // LATIN CAPITAL LETTER F
		'\x47':	'\u0047',	 // LATIN CAPITAL LETTER G
		'\x48':	'\u0048',	 // LATIN CAPITAL LETTER H
		'\x49':	'\u0049',	 // LATIN CAPITAL LETTER I
		'\x4A':	'\u004A',	 // LATIN CAPITAL LETTER J
		'\x4B':	'\u004B',	 // LATIN CAPITAL LETTER K
		'\x4C':	'\u004C',	 // LATIN CAPITAL LETTER L
		'\x4D':	'\u004D',	 // LATIN CAPITAL LETTER M
		'\x4E':	'\u004E',	 // LATIN CAPITAL LETTER N
		'\x4F':	'\u004F',	 // LATIN CAPITAL LETTER O
		'\x50':	'\u0050',	 // LATIN CAPITAL LETTER P
		'\x51':	'\u0051',	 // LATIN CAPITAL LETTER Q
		'\x52':	'\u0052',	 // LATIN CAPITAL LETTER R
		'\x53':	'\u0053',	 // LATIN CAPITAL LETTER S
		'\x54':	'\u0054',	 // LATIN CAPITAL LETTER T
		'\x55':	'\u0055',	 // LATIN CAPITAL LETTER U
		'\x56':	'\u0056',	 // LATIN CAPITAL LETTER V
		'\x57':	'\u0057',	 // LATIN CAPITAL LETTER W
		'\x58':	'\u0058',	 // LATIN CAPITAL LETTER X
		'\x59':	'\u0059',	 // LATIN CAPITAL LETTER Y
		'\x5A':	'\u005A',	 // LATIN CAPITAL LETTER Z
		'\x5B':	'\u005B',	 // LEFT SQUARE BRACKET
		'\x5C':	'\u005C',	 // REVERSE SOLIDUS
		'\x5D':	'\u005D',	 // RIGHT SQUARE BRACKET
		'\x5E':	'\u005E',	 // CIRCUMFLEX ACCENT
		'\x5F':	'\u005F',	 // LOW LINE
		'\x60':	'\u0060',	 // GRAVE ACCENT
		'\x61':	'\u0061',	 // LATIN SMALL LETTER A
		'\x62':	'\u0062',	 // LATIN SMALL LETTER B
		'\x63':	'\u0063',	 // LATIN SMALL LETTER C
		'\x64':	'\u0064',	 // LATIN SMALL LETTER D
		'\x65':	'\u0065',	 // LATIN SMALL LETTER E
		'\x66':	'\u0066',	 // LATIN SMALL LETTER F
		'\x67':	'\u0067',	 // LATIN SMALL LETTER G
		'\x68':	'\u0068',	 // LATIN SMALL LETTER H
		'\x69':	'\u0069',	 // LATIN SMALL LETTER I
		'\x6A':	'\u006A',	 // LATIN SMALL LETTER J
		'\x6B':	'\u006B',	 // LATIN SMALL LETTER K
		'\x6C':	'\u006C',	 // LATIN SMALL LETTER L
		'\x6D':	'\u006D',	 // LATIN SMALL LETTER M
		'\x6E':	'\u006E',	 // LATIN SMALL LETTER N
		'\x6F':	'\u006F',	 // LATIN SMALL LETTER O
		'\x70':	'\u0070',	 // LATIN SMALL LETTER P
		'\x71':	'\u0071',	 // LATIN SMALL LETTER Q
		'\x72':	'\u0072',	 // LATIN SMALL LETTER R
		'\x73':	'\u0073',	 // LATIN SMALL LETTER S
		'\x74':	'\u0074',	 // LATIN SMALL LETTER T
		'\x75':	'\u0075',	 // LATIN SMALL LETTER U
		'\x76':	'\u0076',	 // LATIN SMALL LETTER V
		'\x77':	'\u0077',	 // LATIN SMALL LETTER W
		'\x78':	'\u0078',	 // LATIN SMALL LETTER X
		'\x79':	'\u0079',	 // LATIN SMALL LETTER Y
		'\x7A':	'\u007A',	 // LATIN SMALL LETTER Z
		'\x7B':	'\u007B',	 // LEFT CURLY BRACKET
		'\x7C':	'\u007C',	 // VERTICAL LINE
		'\x7D':	'\u007D',	 // RIGHT CURLY BRACKET
		'\x7E':	'\u007E',	 // TILDE
		'\x7F':	'\u007F',	 // DELETE
		// '\x80' UNDEFINED
		// '\x81' DBCS LEAD BYTE
		// '\x82' DBCS LEAD BYTE
		// '\x83' DBCS LEAD BYTE
		// '\x84' DBCS LEAD BYTE
		// '\x85' DBCS LEAD BYTE
		// '\x86' DBCS LEAD BYTE
		// '\x87' DBCS LEAD BYTE
		// '\x88' DBCS LEAD BYTE
		// '\x89' DBCS LEAD BYTE
		// '\x8A' DBCS LEAD BYTE
		// '\x8B' DBCS LEAD BYTE
		// '\x8C' DBCS LEAD BYTE
		// '\x8D' DBCS LEAD BYTE
		// '\x8E' DBCS LEAD BYTE
		// '\x8F' DBCS LEAD BYTE
		// '\x90' DBCS LEAD BYTE
		// '\x91' DBCS LEAD BYTE
		// '\x92' DBCS LEAD BYTE
		// '\x93' DBCS LEAD BYTE
		// '\x94' DBCS LEAD BYTE
		// '\x95' DBCS LEAD BYTE
		// '\x96' DBCS LEAD BYTE
		// '\x97' DBCS LEAD BYTE
		// '\x98' DBCS LEAD BYTE
		// '\x99' DBCS LEAD BYTE
		// '\x9A' DBCS LEAD BYTE
		// '\x9B' DBCS LEAD BYTE
		// '\x9C' DBCS LEAD BYTE
		// '\x9D' DBCS LEAD BYTE
		// '\x9E' DBCS LEAD BYTE
		// '\x9F' DBCS LEAD BYTE
		// '\xA0' DBCS LEAD BYTE
		// '\xA1' DBCS LEAD BYTE
		// '\xA2' DBCS LEAD BYTE
		// '\xA3' DBCS LEAD BYTE
		// '\xA4' DBCS LEAD BYTE
		// '\xA5' DBCS LEAD BYTE
		// '\xA6' DBCS LEAD BYTE
		// '\xA7' DBCS LEAD BYTE
		// '\xA8' DBCS LEAD BYTE
		// '\xA9' DBCS LEAD BYTE
		// '\xAA' DBCS LEAD BYTE
		// '\xAB' DBCS LEAD BYTE
		// '\xAC' DBCS LEAD BYTE
		// '\xAD' DBCS LEAD BYTE
		// '\xAE' DBCS LEAD BYTE
		// '\xAF' DBCS LEAD BYTE
		// '\xB0' DBCS LEAD BYTE
		// '\xB1' DBCS LEAD BYTE
		// '\xB2' DBCS LEAD BYTE
		// '\xB3' DBCS LEAD BYTE
		// '\xB4' DBCS LEAD BYTE
		// '\xB5' DBCS LEAD BYTE
		// '\xB6' DBCS LEAD BYTE
		// '\xB7' DBCS LEAD BYTE
		// '\xB8' DBCS LEAD BYTE
		// '\xB9' DBCS LEAD BYTE
		// '\xBA' DBCS LEAD BYTE
		// '\xBB' DBCS LEAD BYTE
		// '\xBC' DBCS LEAD BYTE
		// '\xBD' DBCS LEAD BYTE
		// '\xBE' DBCS LEAD BYTE
		// '\xBF' DBCS LEAD BYTE
		// '\xC0' DBCS LEAD BYTE
		// '\xC1' DBCS LEAD BYTE
		// '\xC2' DBCS LEAD BYTE
		// '\xC3' DBCS LEAD BYTE
		// '\xC4' DBCS LEAD BYTE
		// '\xC5' DBCS LEAD BYTE
		// '\xC6' DBCS LEAD BYTE
		// '\xC7' DBCS LEAD BYTE
		// '\xC8' DBCS LEAD BYTE
		// '\xC9' DBCS LEAD BYTE
		// '\xCA' DBCS LEAD BYTE
		// '\xCB' DBCS LEAD BYTE
		// '\xCC' DBCS LEAD BYTE
		// '\xCD' DBCS LEAD BYTE
		// '\xCE' DBCS LEAD BYTE
		// '\xCF' DBCS LEAD BYTE
		// '\xD0' DBCS LEAD BYTE
		// '\xD1' DBCS LEAD BYTE
		// '\xD2' DBCS LEAD BYTE
		// '\xD3' DBCS LEAD BYTE
		// '\xD4' DBCS LEAD BYTE
		// '\xD5' DBCS LEAD BYTE
		// '\xD6' DBCS LEAD BYTE
		// '\xD7' DBCS LEAD BYTE
		// '\xD8' DBCS LEAD BYTE
		// '\xD9' DBCS LEAD BYTE
		// '\xDA' DBCS LEAD BYTE
		// '\xDB' DBCS LEAD BYTE
		// '\xDC' DBCS LEAD BYTE
		// '\xDD' DBCS LEAD BYTE
		// '\xDE' DBCS LEAD BYTE
		// '\xDF' DBCS LEAD BYTE
		// '\xE0' DBCS LEAD BYTE
		// '\xE1' DBCS LEAD BYTE
		// '\xE2' DBCS LEAD BYTE
		// '\xE3' DBCS LEAD BYTE
		// '\xE4' DBCS LEAD BYTE
		// '\xE5' DBCS LEAD BYTE
		// '\xE6' DBCS LEAD BYTE
		// '\xE7' DBCS LEAD BYTE
		// '\xE8' DBCS LEAD BYTE
		// '\xE9' DBCS LEAD BYTE
		// '\xEA' DBCS LEAD BYTE
		// '\xEB' DBCS LEAD BYTE
		// '\xEC' DBCS LEAD BYTE
		// '\xED' DBCS LEAD BYTE
		// '\xEE' DBCS LEAD BYTE
		// '\xEF' DBCS LEAD BYTE
		// '\xF0' DBCS LEAD BYTE
		// '\xF1' DBCS LEAD BYTE
		// '\xF2' DBCS LEAD BYTE
		// '\xF3' DBCS LEAD BYTE
		// '\xF4' DBCS LEAD BYTE
		// '\xF5' DBCS LEAD BYTE
		// '\xF6' DBCS LEAD BYTE
		// '\xF7' DBCS LEAD BYTE
		// '\xF8' DBCS LEAD BYTE
		// '\xF9' DBCS LEAD BYTE
		// '\xFA' DBCS LEAD BYTE
		// '\xFB' DBCS LEAD BYTE
		// '\xFC' DBCS LEAD BYTE
		// '\xFD' DBCS LEAD BYTE
		// '\xFE' DBCS LEAD BYTE
		// '\xFF' UNDEFINED

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecCP949{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "cp949",
		Aliases: []string{
			"949",
		},
		Codec: codec,
	}

	Register(cm)
}
