
package charmap

type codecCP775 struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecCP775) Encode(s string) (string, error) {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecCP775) Decode(s string) (string, error) {
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
		'\x0a':	'\u000a',	 // LINE FEED
		'\x0b':	'\u000b',	 // VERTICAL TABULATION
		'\x0c':	'\u000c',	 // FORM FEED
		'\x0d':	'\u000d',	 // CARRIAGE RETURN
		'\x0e':	'\u000e',	 // SHIFT OUT
		'\x0f':	'\u000f',	 // SHIFT IN
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
		'\x1a':	'\u001a',	 // SUBSTITUTE
		'\x1b':	'\u001b',	 // ESCAPE
		'\x1c':	'\u001c',	 // FILE SEPARATOR
		'\x1d':	'\u001d',	 // GROUP SEPARATOR
		'\x1e':	'\u001e',	 // RECORD SEPARATOR
		'\x1f':	'\u001f',	 // UNIT SEPARATOR
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
		'\x2a':	'\u002a',	 // ASTERISK
		'\x2b':	'\u002b',	 // PLUS SIGN
		'\x2c':	'\u002c',	 // COMMA
		'\x2d':	'\u002d',	 // HYPHEN-MINUS
		'\x2e':	'\u002e',	 // FULL STOP
		'\x2f':	'\u002f',	 // SOLIDUS
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
		'\x3a':	'\u003a',	 // COLON
		'\x3b':	'\u003b',	 // SEMICOLON
		'\x3c':	'\u003c',	 // LESS-THAN SIGN
		'\x3d':	'\u003d',	 // EQUALS SIGN
		'\x3e':	'\u003e',	 // GREATER-THAN SIGN
		'\x3f':	'\u003f',	 // QUESTION MARK
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
		'\x4a':	'\u004a',	 // LATIN CAPITAL LETTER J
		'\x4b':	'\u004b',	 // LATIN CAPITAL LETTER K
		'\x4c':	'\u004c',	 // LATIN CAPITAL LETTER L
		'\x4d':	'\u004d',	 // LATIN CAPITAL LETTER M
		'\x4e':	'\u004e',	 // LATIN CAPITAL LETTER N
		'\x4f':	'\u004f',	 // LATIN CAPITAL LETTER O
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
		'\x5a':	'\u005a',	 // LATIN CAPITAL LETTER Z
		'\x5b':	'\u005b',	 // LEFT SQUARE BRACKET
		'\x5c':	'\u005c',	 // REVERSE SOLIDUS
		'\x5d':	'\u005d',	 // RIGHT SQUARE BRACKET
		'\x5e':	'\u005e',	 // CIRCUMFLEX ACCENT
		'\x5f':	'\u005f',	 // LOW LINE
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
		'\x6a':	'\u006a',	 // LATIN SMALL LETTER J
		'\x6b':	'\u006b',	 // LATIN SMALL LETTER K
		'\x6c':	'\u006c',	 // LATIN SMALL LETTER L
		'\x6d':	'\u006d',	 // LATIN SMALL LETTER M
		'\x6e':	'\u006e',	 // LATIN SMALL LETTER N
		'\x6f':	'\u006f',	 // LATIN SMALL LETTER O
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
		'\x7a':	'\u007a',	 // LATIN SMALL LETTER Z
		'\x7b':	'\u007b',	 // LEFT CURLY BRACKET
		'\x7c':	'\u007c',	 // VERTICAL LINE
		'\x7d':	'\u007d',	 // RIGHT CURLY BRACKET
		'\x7e':	'\u007e',	 // TILDE
		'\x7f':	'\u007f',	 // DELETE
		'\x80':	'\u0106',	 // LATIN CAPITAL LETTER C WITH ACUTE
		'\x81':	'\u00fc',	 // LATIN SMALL LETTER U WITH DIAERESIS
		'\x82':	'\u00e9',	 // LATIN SMALL LETTER E WITH ACUTE
		'\x83':	'\u0101',	 // LATIN SMALL LETTER A WITH MACRON
		'\x84':	'\u00e4',	 // LATIN SMALL LETTER A WITH DIAERESIS
		'\x85':	'\u0123',	 // LATIN SMALL LETTER G WITH CEDILLA
		'\x86':	'\u00e5',	 // LATIN SMALL LETTER A WITH RING ABOVE
		'\x87':	'\u0107',	 // LATIN SMALL LETTER C WITH ACUTE
		'\x88':	'\u0142',	 // LATIN SMALL LETTER L WITH STROKE
		'\x89':	'\u0113',	 // LATIN SMALL LETTER E WITH MACRON
		'\x8a':	'\u0156',	 // LATIN CAPITAL LETTER R WITH CEDILLA
		'\x8b':	'\u0157',	 // LATIN SMALL LETTER R WITH CEDILLA
		'\x8c':	'\u012b',	 // LATIN SMALL LETTER I WITH MACRON
		'\x8d':	'\u0179',	 // LATIN CAPITAL LETTER Z WITH ACUTE
		'\x8e':	'\u00c4',	 // LATIN CAPITAL LETTER A WITH DIAERESIS
		'\x8f':	'\u00c5',	 // LATIN CAPITAL LETTER A WITH RING ABOVE
		'\x90':	'\u00c9',	 // LATIN CAPITAL LETTER E WITH ACUTE
		'\x91':	'\u00e6',	 // LATIN SMALL LIGATURE AE
		'\x92':	'\u00c6',	 // LATIN CAPITAL LIGATURE AE
		'\x93':	'\u014d',	 // LATIN SMALL LETTER O WITH MACRON
		'\x94':	'\u00f6',	 // LATIN SMALL LETTER O WITH DIAERESIS
		'\x95':	'\u0122',	 // LATIN CAPITAL LETTER G WITH CEDILLA
		'\x96':	'\u00a2',	 // CENT SIGN
		'\x97':	'\u015a',	 // LATIN CAPITAL LETTER S WITH ACUTE
		'\x98':	'\u015b',	 // LATIN SMALL LETTER S WITH ACUTE
		'\x99':	'\u00d6',	 // LATIN CAPITAL LETTER O WITH DIAERESIS
		'\x9a':	'\u00dc',	 // LATIN CAPITAL LETTER U WITH DIAERESIS
		'\x9b':	'\u00f8',	 // LATIN SMALL LETTER O WITH STROKE
		'\x9c':	'\u00a3',	 // POUND SIGN
		'\x9d':	'\u00d8',	 // LATIN CAPITAL LETTER O WITH STROKE
		'\x9e':	'\u00d7',	 // MULTIPLICATION SIGN
		'\x9f':	'\u00a4',	 // CURRENCY SIGN
		'\xa0':	'\u0100',	 // LATIN CAPITAL LETTER A WITH MACRON
		'\xa1':	'\u012a',	 // LATIN CAPITAL LETTER I WITH MACRON
		'\xa2':	'\u00f3',	 // LATIN SMALL LETTER O WITH ACUTE
		'\xa3':	'\u017b',	 // LATIN CAPITAL LETTER Z WITH DOT ABOVE
		'\xa4':	'\u017c',	 // LATIN SMALL LETTER Z WITH DOT ABOVE
		'\xa5':	'\u017a',	 // LATIN SMALL LETTER Z WITH ACUTE
		'\xa6':	'\u201d',	 // RIGHT DOUBLE QUOTATION MARK
		'\xa7':	'\u00a6',	 // BROKEN BAR
		'\xa8':	'\u00a9',	 // COPYRIGHT SIGN
		'\xa9':	'\u00ae',	 // REGISTERED SIGN
		'\xaa':	'\u00ac',	 // NOT SIGN
		'\xab':	'\u00bd',	 // VULGAR FRACTION ONE HALF
		'\xac':	'\u00bc',	 // VULGAR FRACTION ONE QUARTER
		'\xad':	'\u0141',	 // LATIN CAPITAL LETTER L WITH STROKE
		'\xae':	'\u00ab',	 // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xaf':	'\u00bb',	 // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xb0':	'\u2591',	 // LIGHT SHADE
		'\xb1':	'\u2592',	 // MEDIUM SHADE
		'\xb2':	'\u2593',	 // DARK SHADE
		'\xb3':	'\u2502',	 // BOX DRAWINGS LIGHT VERTICAL
		'\xb4':	'\u2524',	 // BOX DRAWINGS LIGHT VERTICAL AND LEFT
		'\xb5':	'\u0104',	 // LATIN CAPITAL LETTER A WITH OGONEK
		'\xb6':	'\u010c',	 // LATIN CAPITAL LETTER C WITH CARON
		'\xb7':	'\u0118',	 // LATIN CAPITAL LETTER E WITH OGONEK
		'\xb8':	'\u0116',	 // LATIN CAPITAL LETTER E WITH DOT ABOVE
		'\xb9':	'\u2563',	 // BOX DRAWINGS DOUBLE VERTICAL AND LEFT
		'\xba':	'\u2551',	 // BOX DRAWINGS DOUBLE VERTICAL
		'\xbb':	'\u2557',	 // BOX DRAWINGS DOUBLE DOWN AND LEFT
		'\xbc':	'\u255d',	 // BOX DRAWINGS DOUBLE UP AND LEFT
		'\xbd':	'\u012e',	 // LATIN CAPITAL LETTER I WITH OGONEK
		'\xbe':	'\u0160',	 // LATIN CAPITAL LETTER S WITH CARON
		'\xbf':	'\u2510',	 // BOX DRAWINGS LIGHT DOWN AND LEFT
		'\xc0':	'\u2514',	 // BOX DRAWINGS LIGHT UP AND RIGHT
		'\xc1':	'\u2534',	 // BOX DRAWINGS LIGHT UP AND HORIZONTAL
		'\xc2':	'\u252c',	 // BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
		'\xc3':	'\u251c',	 // BOX DRAWINGS LIGHT VERTICAL AND RIGHT
		'\xc4':	'\u2500',	 // BOX DRAWINGS LIGHT HORIZONTAL
		'\xc5':	'\u253c',	 // BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
		'\xc6':	'\u0172',	 // LATIN CAPITAL LETTER U WITH OGONEK
		'\xc7':	'\u016a',	 // LATIN CAPITAL LETTER U WITH MACRON
		'\xc8':	'\u255a',	 // BOX DRAWINGS DOUBLE UP AND RIGHT
		'\xc9':	'\u2554',	 // BOX DRAWINGS DOUBLE DOWN AND RIGHT
		'\xca':	'\u2569',	 // BOX DRAWINGS DOUBLE UP AND HORIZONTAL
		'\xcb':	'\u2566',	 // BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL
		'\xcc':	'\u2560',	 // BOX DRAWINGS DOUBLE VERTICAL AND RIGHT
		'\xcd':	'\u2550',	 // BOX DRAWINGS DOUBLE HORIZONTAL
		'\xce':	'\u256c',	 // BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL
		'\xcf':	'\u017d',	 // LATIN CAPITAL LETTER Z WITH CARON
		'\xd0':	'\u0105',	 // LATIN SMALL LETTER A WITH OGONEK
		'\xd1':	'\u010d',	 // LATIN SMALL LETTER C WITH CARON
		'\xd2':	'\u0119',	 // LATIN SMALL LETTER E WITH OGONEK
		'\xd3':	'\u0117',	 // LATIN SMALL LETTER E WITH DOT ABOVE
		'\xd4':	'\u012f',	 // LATIN SMALL LETTER I WITH OGONEK
		'\xd5':	'\u0161',	 // LATIN SMALL LETTER S WITH CARON
		'\xd6':	'\u0173',	 // LATIN SMALL LETTER U WITH OGONEK
		'\xd7':	'\u016b',	 // LATIN SMALL LETTER U WITH MACRON
		'\xd8':	'\u017e',	 // LATIN SMALL LETTER Z WITH CARON
		'\xd9':	'\u2518',	 // BOX DRAWINGS LIGHT UP AND LEFT
		'\xda':	'\u250c',	 // BOX DRAWINGS LIGHT DOWN AND RIGHT
		'\xdb':	'\u2588',	 // FULL BLOCK
		'\xdc':	'\u2584',	 // LOWER HALF BLOCK
		'\xdd':	'\u258c',	 // LEFT HALF BLOCK
		'\xde':	'\u2590',	 // RIGHT HALF BLOCK
		'\xdf':	'\u2580',	 // UPPER HALF BLOCK
		'\xe0':	'\u00d3',	 // LATIN CAPITAL LETTER O WITH ACUTE
		'\xe1':	'\u00df',	 // LATIN SMALL LETTER SHARP S (GERMAN)
		'\xe2':	'\u014c',	 // LATIN CAPITAL LETTER O WITH MACRON
		'\xe3':	'\u0143',	 // LATIN CAPITAL LETTER N WITH ACUTE
		'\xe4':	'\u00f5',	 // LATIN SMALL LETTER O WITH TILDE
		'\xe5':	'\u00d5',	 // LATIN CAPITAL LETTER O WITH TILDE
		'\xe6':	'\u00b5',	 // MICRO SIGN
		'\xe7':	'\u0144',	 // LATIN SMALL LETTER N WITH ACUTE
		'\xe8':	'\u0136',	 // LATIN CAPITAL LETTER K WITH CEDILLA
		'\xe9':	'\u0137',	 // LATIN SMALL LETTER K WITH CEDILLA
		'\xea':	'\u013b',	 // LATIN CAPITAL LETTER L WITH CEDILLA
		'\xeb':	'\u013c',	 // LATIN SMALL LETTER L WITH CEDILLA
		'\xec':	'\u0146',	 // LATIN SMALL LETTER N WITH CEDILLA
		'\xed':	'\u0112',	 // LATIN CAPITAL LETTER E WITH MACRON
		'\xee':	'\u0145',	 // LATIN CAPITAL LETTER N WITH CEDILLA
		'\xef':	'\u2019',	 // RIGHT SINGLE QUOTATION MARK
		'\xf0':	'\u00ad',	 // SOFT HYPHEN
		'\xf1':	'\u00b1',	 // PLUS-MINUS SIGN
		'\xf2':	'\u201c',	 // LEFT DOUBLE QUOTATION MARK
		'\xf3':	'\u00be',	 // VULGAR FRACTION THREE QUARTERS
		'\xf4':	'\u00b6',	 // PILCROW SIGN
		'\xf5':	'\u00a7',	 // SECTION SIGN
		'\xf6':	'\u00f7',	 // DIVISION SIGN
		'\xf7':	'\u201e',	 // DOUBLE LOW-9 QUOTATION MARK
		'\xf8':	'\u00b0',	 // DEGREE SIGN
		'\xf9':	'\u2219',	 // BULLET OPERATOR
		'\xfa':	'\u00b7',	 // MIDDLE DOT
		'\xfb':	'\u00b9',	 // SUPERSCRIPT ONE
		'\xfc':	'\u00b3',	 // SUPERSCRIPT THREE
		'\xfd':	'\u00b2',	 // SUPERSCRIPT TWO
		'\xfe':	'\u25a0',	 // BLACK SQUARE
		'\xff':	'\u00a0',	 // NO-BREAK SPACE

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecCP775{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "cp775",
		Aliases: []string{
			"775",
		},
		Codec: codec,
	}

	Register(cm)
}
