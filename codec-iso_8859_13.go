
package charmap

type codecISO_8859_13 struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecISO_8859_13) Encode(s string) (string, error) {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecISO_8859_13) Decode(s string) (string, error) {
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
		'\xA1':	'\u201D',	 // 	RIGHT DOUBLE QUOTATION MARK
		'\xA2':	'\u00A2',	 // 	CENT SIGN
		'\xA3':	'\u00A3',	 // 	POUND SIGN
		'\xA4':	'\u00A4',	 // 	CURRENCY SIGN
		'\xA5':	'\u201E',	 // 	DOUBLE LOW-9 QUOTATION MARK
		'\xA6':	'\u00A6',	 // 	BROKEN BAR
		'\xA7':	'\u00A7',	 // 	SECTION SIGN
		'\xA8':	'\u00D8',	 // 	LATIN CAPITAL LETTER O WITH STROKE
		'\xA9':	'\u00A9',	 // 	COPYRIGHT SIGN
		'\xAA':	'\u0156',	 // 	LATIN CAPITAL LETTER R WITH CEDILLA
		'\xAB':	'\u00AB',	 // 	LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xAC':	'\u00AC',	 // 	NOT SIGN
		'\xAD':	'\u00AD',	 // 	SOFT HYPHEN
		'\xAE':	'\u00AE',	 // 	REGISTERED SIGN
		'\xAF':	'\u00C6',	 // 	LATIN CAPITAL LETTER AE
		'\xB0':	'\u00B0',	 // 	DEGREE SIGN
		'\xB1':	'\u00B1',	 // 	PLUS-MINUS SIGN
		'\xB2':	'\u00B2',	 // 	SUPERSCRIPT TWO
		'\xB3':	'\u00B3',	 // 	SUPERSCRIPT THREE
		'\xB4':	'\u201C',	 // 	LEFT DOUBLE QUOTATION MARK
		'\xB5':	'\u00B5',	 // 	MICRO SIGN
		'\xB6':	'\u00B6',	 // 	PILCROW SIGN
		'\xB7':	'\u00B7',	 // 	MIDDLE DOT
		'\xB8':	'\u00F8',	 // 	LATIN SMALL LETTER O WITH STROKE
		'\xB9':	'\u00B9',	 // 	SUPERSCRIPT ONE
		'\xBA':	'\u0157',	 // 	LATIN SMALL LETTER R WITH CEDILLA
		'\xBB':	'\u00BB',	 // 	RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xBC':	'\u00BC',	 // 	VULGAR FRACTION ONE QUARTER
		'\xBD':	'\u00BD',	 // 	VULGAR FRACTION ONE HALF
		'\xBE':	'\u00BE',	 // 	VULGAR FRACTION THREE QUARTERS
		'\xBF':	'\u00E6',	 // 	LATIN SMALL LETTER AE
		'\xC0':	'\u0104',	 // 	LATIN CAPITAL LETTER A WITH OGONEK
		'\xC1':	'\u012E',	 // 	LATIN CAPITAL LETTER I WITH OGONEK
		'\xC2':	'\u0100',	 // 	LATIN CAPITAL LETTER A WITH MACRON
		'\xC3':	'\u0106',	 // 	LATIN CAPITAL LETTER C WITH ACUTE
		'\xC4':	'\u00C4',	 // 	LATIN CAPITAL LETTER A WITH DIAERESIS
		'\xC5':	'\u00C5',	 // 	LATIN CAPITAL LETTER A WITH RING ABOVE
		'\xC6':	'\u0118',	 // 	LATIN CAPITAL LETTER E WITH OGONEK
		'\xC7':	'\u0112',	 // 	LATIN CAPITAL LETTER E WITH MACRON
		'\xC8':	'\u010C',	 // 	LATIN CAPITAL LETTER C WITH CARON
		'\xC9':	'\u00C9',	 // 	LATIN CAPITAL LETTER E WITH ACUTE
		'\xCA':	'\u0179',	 // 	LATIN CAPITAL LETTER Z WITH ACUTE
		'\xCB':	'\u0116',	 // 	LATIN CAPITAL LETTER E WITH DOT ABOVE
		'\xCC':	'\u0122',	 // 	LATIN CAPITAL LETTER G WITH CEDILLA
		'\xCD':	'\u0136',	 // 	LATIN CAPITAL LETTER K WITH CEDILLA
		'\xCE':	'\u012A',	 // 	LATIN CAPITAL LETTER I WITH MACRON
		'\xCF':	'\u013B',	 // 	LATIN CAPITAL LETTER L WITH CEDILLA
		'\xD0':	'\u0160',	 // 	LATIN CAPITAL LETTER S WITH CARON
		'\xD1':	'\u0143',	 // 	LATIN CAPITAL LETTER N WITH ACUTE
		'\xD2':	'\u0145',	 // 	LATIN CAPITAL LETTER N WITH CEDILLA
		'\xD3':	'\u00D3',	 // 	LATIN CAPITAL LETTER O WITH ACUTE
		'\xD4':	'\u014C',	 // 	LATIN CAPITAL LETTER O WITH MACRON
		'\xD5':	'\u00D5',	 // 	LATIN CAPITAL LETTER O WITH TILDE
		'\xD6':	'\u00D6',	 // 	LATIN CAPITAL LETTER O WITH DIAERESIS
		'\xD7':	'\u00D7',	 // 	MULTIPLICATION SIGN
		'\xD8':	'\u0172',	 // 	LATIN CAPITAL LETTER U WITH OGONEK
		'\xD9':	'\u0141',	 // 	LATIN CAPITAL LETTER L WITH STROKE
		'\xDA':	'\u015A',	 // 	LATIN CAPITAL LETTER S WITH ACUTE
		'\xDB':	'\u016A',	 // 	LATIN CAPITAL LETTER U WITH MACRON
		'\xDC':	'\u00DC',	 // 	LATIN CAPITAL LETTER U WITH DIAERESIS
		'\xDD':	'\u017B',	 // 	LATIN CAPITAL LETTER Z WITH DOT ABOVE
		'\xDE':	'\u017D',	 // 	LATIN CAPITAL LETTER Z WITH CARON
		'\xDF':	'\u00DF',	 // 	LATIN SMALL LETTER SHARP S (German)
		'\xE0':	'\u0105',	 // 	LATIN SMALL LETTER A WITH OGONEK
		'\xE1':	'\u012F',	 // 	LATIN SMALL LETTER I WITH OGONEK
		'\xE2':	'\u0101',	 // 	LATIN SMALL LETTER A WITH MACRON
		'\xE3':	'\u0107',	 // 	LATIN SMALL LETTER C WITH ACUTE
		'\xE4':	'\u00E4',	 // 	LATIN SMALL LETTER A WITH DIAERESIS
		'\xE5':	'\u00E5',	 // 	LATIN SMALL LETTER A WITH RING ABOVE
		'\xE6':	'\u0119',	 // 	LATIN SMALL LETTER E WITH OGONEK
		'\xE7':	'\u0113',	 // 	LATIN SMALL LETTER E WITH MACRON
		'\xE8':	'\u010D',	 // 	LATIN SMALL LETTER C WITH CARON
		'\xE9':	'\u00E9',	 // 	LATIN SMALL LETTER E WITH ACUTE
		'\xEA':	'\u017A',	 // 	LATIN SMALL LETTER Z WITH ACUTE
		'\xEB':	'\u0117',	 // 	LATIN SMALL LETTER E WITH DOT ABOVE
		'\xEC':	'\u0123',	 // 	LATIN SMALL LETTER G WITH CEDILLA
		'\xED':	'\u0137',	 // 	LATIN SMALL LETTER K WITH CEDILLA
		'\xEE':	'\u012B',	 // 	LATIN SMALL LETTER I WITH MACRON
		'\xEF':	'\u013C',	 // 	LATIN SMALL LETTER L WITH CEDILLA
		'\xF0':	'\u0161',	 // 	LATIN SMALL LETTER S WITH CARON
		'\xF1':	'\u0144',	 // 	LATIN SMALL LETTER N WITH ACUTE
		'\xF2':	'\u0146',	 // 	LATIN SMALL LETTER N WITH CEDILLA
		'\xF3':	'\u00F3',	 // 	LATIN SMALL LETTER O WITH ACUTE
		'\xF4':	'\u014D',	 // 	LATIN SMALL LETTER O WITH MACRON
		'\xF5':	'\u00F5',	 // 	LATIN SMALL LETTER O WITH TILDE
		'\xF6':	'\u00F6',	 // 	LATIN SMALL LETTER O WITH DIAERESIS
		'\xF7':	'\u00F7',	 // 	DIVISION SIGN
		'\xF8':	'\u0173',	 // 	LATIN SMALL LETTER U WITH OGONEK
		'\xF9':	'\u0142',	 // 	LATIN SMALL LETTER L WITH STROKE
		'\xFA':	'\u015B',	 // 	LATIN SMALL LETTER S WITH ACUTE
		'\xFB':	'\u016B',	 // 	LATIN SMALL LETTER U WITH MACRON
		'\xFC':	'\u00FC',	 // 	LATIN SMALL LETTER U WITH DIAERESIS
		'\xFD':	'\u017C',	 // 	LATIN SMALL LETTER Z WITH DOT ABOVE
		'\xFE':	'\u017E',	 // 	LATIN SMALL LETTER Z WITH CARON
		'\xFF':	'\u2019',	 // 	RIGHT SINGLE QUOTATION MARK

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecISO_8859_13{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "iso_8859_13",
		Aliases: []string{
			"8859_13",
			"iso8859_13",
		},
		Codec: codec,
	}

	Register(cm)
}
