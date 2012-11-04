
package charmap

type codecISO_8859_5 struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecISO_8859_5) Encode(s string) (string, error) {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecISO_8859_5) Decode(s string) (string, error) {
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
		'\x80':	'\u0080',	 // 	<control>
		'\x81':	'\u0081',	 // 	<control>
		'\x82':	'\u0082',	 // 	<control>
		'\x83':	'\u0083',	 // 	<control>
		'\x84':	'\u0084',	 // 	<control>
		'\x85':	'\u0085',	 // 	<control>
		'\x86':	'\u0086',	 // 	<control>
		'\x87':	'\u0087',	 // 	<control>
		'\x88':	'\u0088',	 // 	<control>
		'\x89':	'\u0089',	 // 	<control>
		'\x8A':	'\u008A',	 // 	<control>
		'\x8B':	'\u008B',	 // 	<control>
		'\x8C':	'\u008C',	 // 	<control>
		'\x8D':	'\u008D',	 // 	<control>
		'\x8E':	'\u008E',	 // 	<control>
		'\x8F':	'\u008F',	 // 	<control>
		'\x90':	'\u0090',	 // 	<control>
		'\x91':	'\u0091',	 // 	<control>
		'\x92':	'\u0092',	 // 	<control>
		'\x93':	'\u0093',	 // 	<control>
		'\x94':	'\u0094',	 // 	<control>
		'\x95':	'\u0095',	 // 	<control>
		'\x96':	'\u0096',	 // 	<control>
		'\x97':	'\u0097',	 // 	<control>
		'\x98':	'\u0098',	 // 	<control>
		'\x99':	'\u0099',	 // 	<control>
		'\x9A':	'\u009A',	 // 	<control>
		'\x9B':	'\u009B',	 // 	<control>
		'\x9C':	'\u009C',	 // 	<control>
		'\x9D':	'\u009D',	 // 	<control>
		'\x9E':	'\u009E',	 // 	<control>
		'\x9F':	'\u009F',	 // 	<control>
		'\xA0':	'\u00A0',	 // 	NO-BREAK SPACE
		'\xA1':	'\u0401',	 // 	CYRILLIC CAPITAL LETTER IO
		'\xA2':	'\u0402',	 // 	CYRILLIC CAPITAL LETTER DJE
		'\xA3':	'\u0403',	 // 	CYRILLIC CAPITAL LETTER GJE
		'\xA4':	'\u0404',	 // 	CYRILLIC CAPITAL LETTER UKRAINIAN IE
		'\xA5':	'\u0405',	 // 	CYRILLIC CAPITAL LETTER DZE
		'\xA6':	'\u0406',	 // 	CYRILLIC CAPITAL LETTER BYELORUSSIAN-UKRAINIAN I
		'\xA7':	'\u0407',	 // 	CYRILLIC CAPITAL LETTER YI
		'\xA8':	'\u0408',	 // 	CYRILLIC CAPITAL LETTER JE
		'\xA9':	'\u0409',	 // 	CYRILLIC CAPITAL LETTER LJE
		'\xAA':	'\u040A',	 // 	CYRILLIC CAPITAL LETTER NJE
		'\xAB':	'\u040B',	 // 	CYRILLIC CAPITAL LETTER TSHE
		'\xAC':	'\u040C',	 // 	CYRILLIC CAPITAL LETTER KJE
		'\xAD':	'\u00AD',	 // 	SOFT HYPHEN
		'\xAE':	'\u040E',	 // 	CYRILLIC CAPITAL LETTER SHORT U
		'\xAF':	'\u040F',	 // 	CYRILLIC CAPITAL LETTER DZHE
		'\xB0':	'\u0410',	 // 	CYRILLIC CAPITAL LETTER A
		'\xB1':	'\u0411',	 // 	CYRILLIC CAPITAL LETTER BE
		'\xB2':	'\u0412',	 // 	CYRILLIC CAPITAL LETTER VE
		'\xB3':	'\u0413',	 // 	CYRILLIC CAPITAL LETTER GHE
		'\xB4':	'\u0414',	 // 	CYRILLIC CAPITAL LETTER DE
		'\xB5':	'\u0415',	 // 	CYRILLIC CAPITAL LETTER IE
		'\xB6':	'\u0416',	 // 	CYRILLIC CAPITAL LETTER ZHE
		'\xB7':	'\u0417',	 // 	CYRILLIC CAPITAL LETTER ZE
		'\xB8':	'\u0418',	 // 	CYRILLIC CAPITAL LETTER I
		'\xB9':	'\u0419',	 // 	CYRILLIC CAPITAL LETTER SHORT I
		'\xBA':	'\u041A',	 // 	CYRILLIC CAPITAL LETTER KA
		'\xBB':	'\u041B',	 // 	CYRILLIC CAPITAL LETTER EL
		'\xBC':	'\u041C',	 // 	CYRILLIC CAPITAL LETTER EM
		'\xBD':	'\u041D',	 // 	CYRILLIC CAPITAL LETTER EN
		'\xBE':	'\u041E',	 // 	CYRILLIC CAPITAL LETTER O
		'\xBF':	'\u041F',	 // 	CYRILLIC CAPITAL LETTER PE
		'\xC0':	'\u0420',	 // 	CYRILLIC CAPITAL LETTER ER
		'\xC1':	'\u0421',	 // 	CYRILLIC CAPITAL LETTER ES
		'\xC2':	'\u0422',	 // 	CYRILLIC CAPITAL LETTER TE
		'\xC3':	'\u0423',	 // 	CYRILLIC CAPITAL LETTER U
		'\xC4':	'\u0424',	 // 	CYRILLIC CAPITAL LETTER EF
		'\xC5':	'\u0425',	 // 	CYRILLIC CAPITAL LETTER HA
		'\xC6':	'\u0426',	 // 	CYRILLIC CAPITAL LETTER TSE
		'\xC7':	'\u0427',	 // 	CYRILLIC CAPITAL LETTER CHE
		'\xC8':	'\u0428',	 // 	CYRILLIC CAPITAL LETTER SHA
		'\xC9':	'\u0429',	 // 	CYRILLIC CAPITAL LETTER SHCHA
		'\xCA':	'\u042A',	 // 	CYRILLIC CAPITAL LETTER HARD SIGN
		'\xCB':	'\u042B',	 // 	CYRILLIC CAPITAL LETTER YERU
		'\xCC':	'\u042C',	 // 	CYRILLIC CAPITAL LETTER SOFT SIGN
		'\xCD':	'\u042D',	 // 	CYRILLIC CAPITAL LETTER E
		'\xCE':	'\u042E',	 // 	CYRILLIC CAPITAL LETTER YU
		'\xCF':	'\u042F',	 // 	CYRILLIC CAPITAL LETTER YA
		'\xD0':	'\u0430',	 // 	CYRILLIC SMALL LETTER A
		'\xD1':	'\u0431',	 // 	CYRILLIC SMALL LETTER BE
		'\xD2':	'\u0432',	 // 	CYRILLIC SMALL LETTER VE
		'\xD3':	'\u0433',	 // 	CYRILLIC SMALL LETTER GHE
		'\xD4':	'\u0434',	 // 	CYRILLIC SMALL LETTER DE
		'\xD5':	'\u0435',	 // 	CYRILLIC SMALL LETTER IE
		'\xD6':	'\u0436',	 // 	CYRILLIC SMALL LETTER ZHE
		'\xD7':	'\u0437',	 // 	CYRILLIC SMALL LETTER ZE
		'\xD8':	'\u0438',	 // 	CYRILLIC SMALL LETTER I
		'\xD9':	'\u0439',	 // 	CYRILLIC SMALL LETTER SHORT I
		'\xDA':	'\u043A',	 // 	CYRILLIC SMALL LETTER KA
		'\xDB':	'\u043B',	 // 	CYRILLIC SMALL LETTER EL
		'\xDC':	'\u043C',	 // 	CYRILLIC SMALL LETTER EM
		'\xDD':	'\u043D',	 // 	CYRILLIC SMALL LETTER EN
		'\xDE':	'\u043E',	 // 	CYRILLIC SMALL LETTER O
		'\xDF':	'\u043F',	 // 	CYRILLIC SMALL LETTER PE
		'\xE0':	'\u0440',	 // 	CYRILLIC SMALL LETTER ER
		'\xE1':	'\u0441',	 // 	CYRILLIC SMALL LETTER ES
		'\xE2':	'\u0442',	 // 	CYRILLIC SMALL LETTER TE
		'\xE3':	'\u0443',	 // 	CYRILLIC SMALL LETTER U
		'\xE4':	'\u0444',	 // 	CYRILLIC SMALL LETTER EF
		'\xE5':	'\u0445',	 // 	CYRILLIC SMALL LETTER HA
		'\xE6':	'\u0446',	 // 	CYRILLIC SMALL LETTER TSE
		'\xE7':	'\u0447',	 // 	CYRILLIC SMALL LETTER CHE
		'\xE8':	'\u0448',	 // 	CYRILLIC SMALL LETTER SHA
		'\xE9':	'\u0449',	 // 	CYRILLIC SMALL LETTER SHCHA
		'\xEA':	'\u044A',	 // 	CYRILLIC SMALL LETTER HARD SIGN
		'\xEB':	'\u044B',	 // 	CYRILLIC SMALL LETTER YERU
		'\xEC':	'\u044C',	 // 	CYRILLIC SMALL LETTER SOFT SIGN
		'\xED':	'\u044D',	 // 	CYRILLIC SMALL LETTER E
		'\xEE':	'\u044E',	 // 	CYRILLIC SMALL LETTER YU
		'\xEF':	'\u044F',	 // 	CYRILLIC SMALL LETTER YA
		'\xF0':	'\u2116',	 // 	NUMERO SIGN
		'\xF1':	'\u0451',	 // 	CYRILLIC SMALL LETTER IO
		'\xF2':	'\u0452',	 // 	CYRILLIC SMALL LETTER DJE
		'\xF3':	'\u0453',	 // 	CYRILLIC SMALL LETTER GJE
		'\xF4':	'\u0454',	 // 	CYRILLIC SMALL LETTER UKRAINIAN IE
		'\xF5':	'\u0455',	 // 	CYRILLIC SMALL LETTER DZE
		'\xF6':	'\u0456',	 // 	CYRILLIC SMALL LETTER BYELORUSSIAN-UKRAINIAN I
		'\xF7':	'\u0457',	 // 	CYRILLIC SMALL LETTER YI
		'\xF8':	'\u0458',	 // 	CYRILLIC SMALL LETTER JE
		'\xF9':	'\u0459',	 // 	CYRILLIC SMALL LETTER LJE
		'\xFA':	'\u045A',	 // 	CYRILLIC SMALL LETTER NJE
		'\xFB':	'\u045B',	 // 	CYRILLIC SMALL LETTER TSHE
		'\xFC':	'\u045C',	 // 	CYRILLIC SMALL LETTER KJE
		'\xFD':	'\u00A7',	 // 	SECTION SIGN
		'\xFE':	'\u045E',	 // 	CYRILLIC SMALL LETTER SHORT U
		'\xFF':	'\u045F',	 // 	CYRILLIC SMALL LETTER DZHE

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecISO_8859_5{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "iso_8859_5",
		Aliases: []string{
			"8859_5",
			"iso8859_5",
		},
		Codec: codec,
	}

	Register(cm)
}
