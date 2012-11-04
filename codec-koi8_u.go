
package charmap

type codecKOI8_U struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecKOI8_U) Encode(s string) (string, error) {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecKOI8_U) Decode(s string) (string, error) {
	return mapBytesToRunes(c.DecodeMap, s)
}

func init() {
	charmapDecode := map[byte]rune{
		'\x00':	'\u0000',	 // 	NULL
		'\x01':	'\u0001',	 // 	START OF HEADING
		'\x02':	'\u0002',	 // 	START OF TEXT
		'\x03':	'\u0003',	 // 	END OF TEXT
		'\x04':	'\u0004',	 // 	END OF TRANSMISSION
		'\x05':	'\u0005',	 // 	ENQUIRY
		'\x06':	'\u0006',	 // 	ACKNOWLEDGE
		'\x07':	'\u0007',	 // 	BELL
		'\x08':	'\u0008',	 // 	BACKSPACE
		'\x09':	'\u0009',	 // 	HORIZONTAL TABULATION
		'\x0A':	'\u000A',	 // 	LINE FEED
		'\x0B':	'\u000B',	 // 	VERTICAL TABULATION
		'\x0C':	'\u000C',	 // 	FORM FEED
		'\x0D':	'\u000D',	 // 	CARRIAGE RETURN
		'\x0E':	'\u000E',	 // 	SHIFT OUT
		'\x0F':	'\u000F',	 // 	SHIFT IN
		'\x10':	'\u0010',	 // 	DATA LINK ESCAPE
		'\x11':	'\u0011',	 // 	DEVICE CONTROL ONE
		'\x12':	'\u0012',	 // 	DEVICE CONTROL TWO
		'\x13':	'\u0013',	 // 	DEVICE CONTROL THREE
		'\x14':	'\u0014',	 // 	DEVICE CONTROL FOUR
		'\x15':	'\u0015',	 // 	NEGATIVE ACKNOWLEDGE
		'\x16':	'\u0016',	 // 	SYNCHRONOUS IDLE
		'\x17':	'\u0017',	 // 	END OF TRANSMISSION BLOCK
		'\x18':	'\u0018',	 // 	CANCEL
		'\x19':	'\u0019',	 // 	END OF MEDIUM
		'\x1A':	'\u001A',	 // 	SUBSTITUTE
		'\x1B':	'\u001B',	 // 	ESCAPE
		'\x1C':	'\u001C',	 // 	FILE SEPARATOR
		'\x1D':	'\u001D',	 // 	GROUP SEPARATOR
		'\x1E':	'\u001E',	 // 	RECORD SEPARATOR
		'\x1F':	'\u001F',	 // 	UNIT SEPARATOR
		'\x20':	'\u0020',	 // 	SPACE
		'\x21':	'\u0021',	 // 	EXCLAMATION MARK
		'\x22':	'\u0022',	 // 	QUOTATION MARK
		'\x23':	'\u0023',	 // 	NUMBER SIGN
		'\x24':	'\u0024',	 // 	DOLLAR SIGN
		'\x25':	'\u0025',	 // 	PERCENT SIGN
		'\x26':	'\u0026',	 // 	AMPERSAND
		'\x27':	'\u0027',	 // 	APOSTROPHE
		'\x28':	'\u0028',	 // 	LEFT PARENTHESIS
		'\x29':	'\u0029',	 // 	RIGHT PARENTHESIS
		'\x2A':	'\u002A',	 // 	ASTERISK
		'\x2B':	'\u002B',	 // 	PLUS SIGN
		'\x2C':	'\u002C',	 // 	COMMA
		'\x2D':	'\u002D',	 // 	HYPHEN-MINUS
		'\x2E':	'\u002E',	 // 	FULL STOP
		'\x2F':	'\u002F',	 // 	SOLIDUS
		'\x30':	'\u0030',	 // 	DIGIT ZERO
		'\x31':	'\u0031',	 // 	DIGIT ONE
		'\x32':	'\u0032',	 // 	DIGIT TWO
		'\x33':	'\u0033',	 // 	DIGIT THREE
		'\x34':	'\u0034',	 // 	DIGIT FOUR
		'\x35':	'\u0035',	 // 	DIGIT FIVE
		'\x36':	'\u0036',	 // 	DIGIT SIX
		'\x37':	'\u0037',	 // 	DIGIT SEVEN
		'\x38':	'\u0038',	 // 	DIGIT EIGHT
		'\x39':	'\u0039',	 // 	DIGIT NINE
		'\x3A':	'\u003A',	 // 	COLON
		'\x3B':	'\u003B',	 // 	SEMICOLON
		'\x3C':	'\u003C',	 // 	LESS-THAN SIGN
		'\x3D':	'\u003D',	 // 	EQUALS SIGN
		'\x3E':	'\u003E',	 // 	GREATER-THAN SIGN
		'\x3F':	'\u003F',	 // 	QUESTION MARK
		'\x40':	'\u0040',	 // 	COMMERCIAL AT
		'\x41':	'\u0041',	 // 	LATIN CAPITAL LETTER A
		'\x42':	'\u0042',	 // 	LATIN CAPITAL LETTER B
		'\x43':	'\u0043',	 // 	LATIN CAPITAL LETTER C
		'\x44':	'\u0044',	 // 	LATIN CAPITAL LETTER D
		'\x45':	'\u0045',	 // 	LATIN CAPITAL LETTER E
		'\x46':	'\u0046',	 // 	LATIN CAPITAL LETTER F
		'\x47':	'\u0047',	 // 	LATIN CAPITAL LETTER G
		'\x48':	'\u0048',	 // 	LATIN CAPITAL LETTER H
		'\x49':	'\u0049',	 // 	LATIN CAPITAL LETTER I
		'\x4A':	'\u004A',	 // 	LATIN CAPITAL LETTER J
		'\x4B':	'\u004B',	 // 	LATIN CAPITAL LETTER K
		'\x4C':	'\u004C',	 // 	LATIN CAPITAL LETTER L
		'\x4D':	'\u004D',	 // 	LATIN CAPITAL LETTER M
		'\x4E':	'\u004E',	 // 	LATIN CAPITAL LETTER N
		'\x4F':	'\u004F',	 // 	LATIN CAPITAL LETTER O
		'\x50':	'\u0050',	 // 	LATIN CAPITAL LETTER P
		'\x51':	'\u0051',	 // 	LATIN CAPITAL LETTER Q
		'\x52':	'\u0052',	 // 	LATIN CAPITAL LETTER R
		'\x53':	'\u0053',	 // 	LATIN CAPITAL LETTER S
		'\x54':	'\u0054',	 // 	LATIN CAPITAL LETTER T
		'\x55':	'\u0055',	 // 	LATIN CAPITAL LETTER U
		'\x56':	'\u0056',	 // 	LATIN CAPITAL LETTER V
		'\x57':	'\u0057',	 // 	LATIN CAPITAL LETTER W
		'\x58':	'\u0058',	 // 	LATIN CAPITAL LETTER X
		'\x59':	'\u0059',	 // 	LATIN CAPITAL LETTER Y
		'\x5A':	'\u005A',	 // 	LATIN CAPITAL LETTER Z
		'\x5B':	'\u005B',	 // 	LEFT SQUARE BRACKET
		'\x5C':	'\u005C',	 // 	REVERSE SOLIDUS
		'\x5D':	'\u005D',	 // 	RIGHT SQUARE BRACKET
		'\x5E':	'\u005E',	 // 	CIRCUMFLEX ACCENT
		'\x5F':	'\u005F',	 // 	LOW LINE
		'\x60':	'\u0060',	 // 	GRAVE ACCENT
		'\x61':	'\u0061',	 // 	LATIN SMALL LETTER A
		'\x62':	'\u0062',	 // 	LATIN SMALL LETTER B
		'\x63':	'\u0063',	 // 	LATIN SMALL LETTER C
		'\x64':	'\u0064',	 // 	LATIN SMALL LETTER D
		'\x65':	'\u0065',	 // 	LATIN SMALL LETTER E
		'\x66':	'\u0066',	 // 	LATIN SMALL LETTER F
		'\x67':	'\u0067',	 // 	LATIN SMALL LETTER G
		'\x68':	'\u0068',	 // 	LATIN SMALL LETTER H
		'\x69':	'\u0069',	 // 	LATIN SMALL LETTER I
		'\x6A':	'\u006A',	 // 	LATIN SMALL LETTER J
		'\x6B':	'\u006B',	 // 	LATIN SMALL LETTER K
		'\x6C':	'\u006C',	 // 	LATIN SMALL LETTER L
		'\x6D':	'\u006D',	 // 	LATIN SMALL LETTER M
		'\x6E':	'\u006E',	 // 	LATIN SMALL LETTER N
		'\x6F':	'\u006F',	 // 	LATIN SMALL LETTER O
		'\x70':	'\u0070',	 // 	LATIN SMALL LETTER P
		'\x71':	'\u0071',	 // 	LATIN SMALL LETTER Q
		'\x72':	'\u0072',	 // 	LATIN SMALL LETTER R
		'\x73':	'\u0073',	 // 	LATIN SMALL LETTER S
		'\x74':	'\u0074',	 // 	LATIN SMALL LETTER T
		'\x75':	'\u0075',	 // 	LATIN SMALL LETTER U
		'\x76':	'\u0076',	 // 	LATIN SMALL LETTER V
		'\x77':	'\u0077',	 // 	LATIN SMALL LETTER W
		'\x78':	'\u0078',	 // 	LATIN SMALL LETTER X
		'\x79':	'\u0079',	 // 	LATIN SMALL LETTER Y
		'\x7A':	'\u007A',	 // 	LATIN SMALL LETTER Z
		'\x7B':	'\u007B',	 // 	LEFT CURLY BRACKET
		'\x7C':	'\u007C',	 // 	VERTICAL LINE
		'\x7D':	'\u007D',	 // 	RIGHT CURLY BRACKET
		'\x7E':	'\u007E',	 // 	TILDE
		'\x7F':	'\u007F',	 // 	DELETE
		'\x80':	'\u2500',	 // 	BOX DRAWINGS LIGHT HORIZONTAL
		'\x81':	'\u2502',	 // 	BOX DRAWINGS LIGHT VERTICAL
		'\x82':	'\u250C',	 // 	BOX DRAWINGS LIGHT DOWN AND RIGHT
		'\x83':	'\u2510',	 // 	BOX DRAWINGS LIGHT DOWN AND LEFT
		'\x84':	'\u2514',	 // 	BOX DRAWINGS LIGHT UP AND RIGHT
		'\x85':	'\u2518',	 // 	BOX DRAWINGS LIGHT UP AND LEFT
		'\x86':	'\u251C',	 // 	BOX DRAWINGS LIGHT VERTICAL AND RIGHT
		'\x87':	'\u2524',	 // 	BOX DRAWINGS LIGHT VERTICAL AND LEFT
		'\x88':	'\u252C',	 // 	BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
		'\x89':	'\u2534',	 // 	BOX DRAWINGS LIGHT UP AND HORIZONTAL
		'\x8A':	'\u253C',	 // 	BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
		'\x8B':	'\u2580',	 // 	UPPER HALF BLOCK
		'\x8C':	'\u2584',	 // 	LOWER HALF BLOCK
		'\x8D':	'\u2588',	 // 	FULL BLOCK
		'\x8E':	'\u258C',	 // 	LEFT HALF BLOCK
		'\x8F':	'\u2590',	 // 	RIGHT HALF BLOCK
		'\x90':	'\u2591',	 // 	LIGHT SHADE
		'\x91':	'\u2592',	 // 	MEDIUM SHADE
		'\x92':	'\u2593',	 // 	DARK SHADE
		'\x93':	'\u2320',	 // 	TOP HALF INTEGRAL
		'\x94':	'\u25A0',	 // 	BLACK SQUARE
		'\x95':	'\u2219',	 // 	BULLET OPERATOR
		'\x96':	'\u221A',	 // 	SQUARE ROOT
		'\x97':	'\u2248',	 // 	ALMOST EQUAL TO
		'\x98':	'\u2264',	 // 	LESS-THAN OR EQUAL TO
		'\x99':	'\u2265',	 // 	GREATER-THAN OR EQUAL TO
		'\x9A':	'\u00A0',	 // 	NO-BREAK SPACE
		'\x9B':	'\u2321',	 // 	BOTTOM HALF INTEGRAL
		'\x9C':	'\u00B0',	 // 	DEGREE SIGN
		'\x9D':	'\u00B2',	 // 	SUPERSCRIPT TWO
		'\x9E':	'\u00B7',	 // 	MIDDLE DOT
		'\x9F':	'\u00F7',	 // 	DIVISION SIGN
		'\xA0':	'\u2550',	 // 	BOX DRAWINGS DOUBLE HORIZONTAL
		'\xA1':	'\u2551',	 // 	BOX DRAWINGS DOUBLE VERTICAL
		'\xA2':	'\u2552',	 // 	BOX DRAWINGS DOWN SINGLE AND RIGHT DOUBLE
		'\xA3':	'\u0451',	 // 	CYRILLIC SMALL LETTER IO
		'\xA4':	'\u0454',	 // 	CYRILLIC SMALL LETTER UKRAINIAN IE
		'\xA5':	'\u2554',	 // 	BOX DRAWINGS DOUBLE DOWN AND RIGHT
		'\xA6':	'\u0456',	 // 	CYRILLIC SMALL LETTER BYELORUSSIAN-UKRAINIAN I
		'\xA7':	'\u0457',	 // 	CYRILLIC SMALL LETTER YI (UKRAINIAN)
		'\xA8':	'\u2557',	 // 	BOX DRAWINGS DOUBLE DOWN AND LEFT
		'\xA9':	'\u2558',	 // 	BOX DRAWINGS UP SINGLE AND RIGHT DOUBLE
		'\xAA':	'\u2559',	 // 	BOX DRAWINGS UP DOUBLE AND RIGHT SINGLE
		'\xAB':	'\u255A',	 // 	BOX DRAWINGS DOUBLE UP AND RIGHT
		'\xAC':	'\u255B',	 // 	BOX DRAWINGS UP SINGLE AND LEFT DOUBLE
		'\xAD':	'\u0491',	 // 	CYRILLIC SMALL LETTER GHE WITH UPTURN
		'\xAE':	'\u255D',	 // 	BOX DRAWINGS DOUBLE UP AND LEFT
		'\xAF':	'\u255E',	 // 	BOX DRAWINGS VERTICAL SINGLE AND RIGHT DOUBLE
		'\xB0':	'\u255F',	 // 	BOX DRAWINGS VERTICAL DOUBLE AND RIGHT SINGLE
		'\xB1':	'\u2560',	 // 	BOX DRAWINGS DOUBLE VERTICAL AND RIGHT
		'\xB2':	'\u2561',	 // 	BOX DRAWINGS VERTICAL SINGLE AND LEFT DOUBLE
		'\xB3':	'\u0401',	 // 	CYRILLIC CAPITAL LETTER IO
		'\xB4':	'\u0404',	 // 	CYRILLIC CAPITAL LETTER UKRAINIAN IE
		'\xB5':	'\u2563',	 // 	BOX DRAWINGS DOUBLE VERTICAL AND LEFT
		'\xB6':	'\u0406',	 // 	CYRILLIC CAPITAL LETTER BYELORUSSIAN-UKRAINIAN I
		'\xB7':	'\u0407',	 // 	CYRILLIC CAPITAL LETTER YI (UKRAINIAN)
		'\xB8':	'\u2566',	 // 	BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL
		'\xB9':	'\u2567',	 // 	BOX DRAWINGS UP SINGLE AND HORIZONTAL DOUBLE
		'\xBA':	'\u2568',	 // 	BOX DRAWINGS UP DOUBLE AND HORIZONTAL SINGLE
		'\xBB':	'\u2569',	 // 	BOX DRAWINGS DOUBLE UP AND HORIZONTAL
		'\xBC':	'\u256A',	 // 	BOX DRAWINGS VERTICAL SINGLE AND HORIZONTAL DOUBLE
		'\xBD':	'\u0490',	 // 	CYRILLIC CAPITAL LETTER GHE WITH UPTURN
		'\xBE':	'\u256C',	 // 	BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL
		'\xBF':	'\u00A9',	 // 	COPYRIGHT SIGN
		'\xC0':	'\u044E',	 // 	CYRILLIC SMALL LETTER YU
		'\xC1':	'\u0430',	 // 	CYRILLIC SMALL LETTER A
		'\xC2':	'\u0431',	 // 	CYRILLIC SMALL LETTER BE
		'\xC3':	'\u0446',	 // 	CYRILLIC SMALL LETTER TSE
		'\xC4':	'\u0434',	 // 	CYRILLIC SMALL LETTER DE
		'\xC5':	'\u0435',	 // 	CYRILLIC SMALL LETTER IE
		'\xC6':	'\u0444',	 // 	CYRILLIC SMALL LETTER EF
		'\xC7':	'\u0433',	 // 	CYRILLIC SMALL LETTER GHE
		'\xC8':	'\u0445',	 // 	CYRILLIC SMALL LETTER HA
		'\xC9':	'\u0438',	 // 	CYRILLIC SMALL LETTER I
		'\xCA':	'\u0439',	 // 	CYRILLIC SMALL LETTER SHORT I
		'\xCB':	'\u043A',	 // 	CYRILLIC SMALL LETTER KA
		'\xCC':	'\u043B',	 // 	CYRILLIC SMALL LETTER EL
		'\xCD':	'\u043C',	 // 	CYRILLIC SMALL LETTER EM
		'\xCE':	'\u043D',	 // 	CYRILLIC SMALL LETTER EN
		'\xCF':	'\u043E',	 // 	CYRILLIC SMALL LETTER O
		'\xD0':	'\u043F',	 // 	CYRILLIC SMALL LETTER PE
		'\xD1':	'\u044F',	 // 	CYRILLIC SMALL LETTER YA
		'\xD2':	'\u0440',	 // 	CYRILLIC SMALL LETTER ER
		'\xD3':	'\u0441',	 // 	CYRILLIC SMALL LETTER ES
		'\xD4':	'\u0442',	 // 	CYRILLIC SMALL LETTER TE
		'\xD5':	'\u0443',	 // 	CYRILLIC SMALL LETTER U
		'\xD6':	'\u0436',	 // 	CYRILLIC SMALL LETTER ZHE
		'\xD7':	'\u0432',	 // 	CYRILLIC SMALL LETTER VE
		'\xD8':	'\u044C',	 // 	CYRILLIC SMALL LETTER SOFT SIGN
		'\xD9':	'\u044B',	 // 	CYRILLIC SMALL LETTER YERU
		'\xDA':	'\u0437',	 // 	CYRILLIC SMALL LETTER ZE
		'\xDB':	'\u0448',	 // 	CYRILLIC SMALL LETTER SHA
		'\xDC':	'\u044D',	 // 	CYRILLIC SMALL LETTER E
		'\xDD':	'\u0449',	 // 	CYRILLIC SMALL LETTER SHCHA
		'\xDE':	'\u0447',	 // 	CYRILLIC SMALL LETTER CHE
		'\xDF':	'\u044A',	 // 	CYRILLIC SMALL LETTER HARD SIGN
		'\xE0':	'\u042E',	 // 	CYRILLIC CAPITAL LETTER YU
		'\xE1':	'\u0410',	 // 	CYRILLIC CAPITAL LETTER A
		'\xE2':	'\u0411',	 // 	CYRILLIC CAPITAL LETTER BE
		'\xE3':	'\u0426',	 // 	CYRILLIC CAPITAL LETTER TSE
		'\xE4':	'\u0414',	 // 	CYRILLIC CAPITAL LETTER DE
		'\xE5':	'\u0415',	 // 	CYRILLIC CAPITAL LETTER IE
		'\xE6':	'\u0424',	 // 	CYRILLIC CAPITAL LETTER EF
		'\xE7':	'\u0413',	 // 	CYRILLIC CAPITAL LETTER GHE
		'\xE8':	'\u0425',	 // 	CYRILLIC CAPITAL LETTER HA
		'\xE9':	'\u0418',	 // 	CYRILLIC CAPITAL LETTER I
		'\xEA':	'\u0419',	 // 	CYRILLIC CAPITAL LETTER SHORT I
		'\xEB':	'\u041A',	 // 	CYRILLIC CAPITAL LETTER KA
		'\xEC':	'\u041B',	 // 	CYRILLIC CAPITAL LETTER EL
		'\xED':	'\u041C',	 // 	CYRILLIC CAPITAL LETTER EM
		'\xEE':	'\u041D',	 // 	CYRILLIC CAPITAL LETTER EN
		'\xEF':	'\u041E',	 // 	CYRILLIC CAPITAL LETTER O
		'\xF0':	'\u041F',	 // 	CYRILLIC CAPITAL LETTER PE
		'\xF1':	'\u042F',	 // 	CYRILLIC CAPITAL LETTER YA
		'\xF2':	'\u0420',	 // 	CYRILLIC CAPITAL LETTER ER
		'\xF3':	'\u0421',	 // 	CYRILLIC CAPITAL LETTER ES
		'\xF4':	'\u0422',	 // 	CYRILLIC CAPITAL LETTER TE
		'\xF5':	'\u0423',	 // 	CYRILLIC CAPITAL LETTER U
		'\xF6':	'\u0416',	 // 	CYRILLIC CAPITAL LETTER ZHE
		'\xF7':	'\u0412',	 // 	CYRILLIC CAPITAL LETTER VE
		'\xF8':	'\u042C',	 // 	CYRILLIC CAPITAL LETTER SOFT SIGN
		'\xF9':	'\u042B',	 // 	CYRILLIC CAPITAL LETTER YERU
		'\xFA':	'\u0417',	 // 	CYRILLIC CAPITAL LETTER ZE
		'\xFB':	'\u0428',	 // 	CYRILLIC CAPITAL LETTER SHA
		'\xFC':	'\u042D',	 // 	CYRILLIC CAPITAL LETTER E
		'\xFD':	'\u0429',	 // 	CYRILLIC CAPITAL LETTER SHCHA
		'\xFE':	'\u0427',	 // 	CYRILLIC CAPITAL LETTER CHE
		'\xFF':	'\u042A',	 // 	CYRILLIC CAPITAL LETTER HARD SIGN

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecKOI8_U{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "koi8_u",
		Aliases: []string{
			"koi8u",
		},
		Codec: codec,
	}

	Register(cm)
}
