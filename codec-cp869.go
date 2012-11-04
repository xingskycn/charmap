
package charmap

type codecCP869 struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecCP869) Encode(s string) (string, error) {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecCP869) Decode(s string) (string, error) {
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
		// '\x80' UNDEFINED
		// '\x81' UNDEFINED
		// '\x82' UNDEFINED
		// '\x83' UNDEFINED
		// '\x84' UNDEFINED
		// '\x85' UNDEFINED
		'\x86':	'\u0386',	 // GREEK CAPITAL LETTER ALPHA WITH TONOS
		// '\x87' UNDEFINED
		'\x88':	'\u00b7',	 // MIDDLE DOT
		'\x89':	'\u00ac',	 // NOT SIGN
		'\x8a':	'\u00a6',	 // BROKEN BAR
		'\x8b':	'\u2018',	 // LEFT SINGLE QUOTATION MARK
		'\x8c':	'\u2019',	 // RIGHT SINGLE QUOTATION MARK
		'\x8d':	'\u0388',	 // GREEK CAPITAL LETTER EPSILON WITH TONOS
		'\x8e':	'\u2015',	 // HORIZONTAL BAR
		'\x8f':	'\u0389',	 // GREEK CAPITAL LETTER ETA WITH TONOS
		'\x90':	'\u038a',	 // GREEK CAPITAL LETTER IOTA WITH TONOS
		'\x91':	'\u03aa',	 // GREEK CAPITAL LETTER IOTA WITH DIALYTIKA
		'\x92':	'\u038c',	 // GREEK CAPITAL LETTER OMICRON WITH TONOS
		// '\x93' UNDEFINED
		// '\x94' UNDEFINED
		'\x95':	'\u038e',	 // GREEK CAPITAL LETTER UPSILON WITH TONOS
		'\x96':	'\u03ab',	 // GREEK CAPITAL LETTER UPSILON WITH DIALYTIKA
		'\x97':	'\u00a9',	 // COPYRIGHT SIGN
		'\x98':	'\u038f',	 // GREEK CAPITAL LETTER OMEGA WITH TONOS
		'\x99':	'\u00b2',	 // SUPERSCRIPT TWO
		'\x9a':	'\u00b3',	 // SUPERSCRIPT THREE
		'\x9b':	'\u03ac',	 // GREEK SMALL LETTER ALPHA WITH TONOS
		'\x9c':	'\u00a3',	 // POUND SIGN
		'\x9d':	'\u03ad',	 // GREEK SMALL LETTER EPSILON WITH TONOS
		'\x9e':	'\u03ae',	 // GREEK SMALL LETTER ETA WITH TONOS
		'\x9f':	'\u03af',	 // GREEK SMALL LETTER IOTA WITH TONOS
		'\xa0':	'\u03ca',	 // GREEK SMALL LETTER IOTA WITH DIALYTIKA
		'\xa1':	'\u0390',	 // GREEK SMALL LETTER IOTA WITH DIALYTIKA AND TONOS
		'\xa2':	'\u03cc',	 // GREEK SMALL LETTER OMICRON WITH TONOS
		'\xa3':	'\u03cd',	 // GREEK SMALL LETTER UPSILON WITH TONOS
		'\xa4':	'\u0391',	 // GREEK CAPITAL LETTER ALPHA
		'\xa5':	'\u0392',	 // GREEK CAPITAL LETTER BETA
		'\xa6':	'\u0393',	 // GREEK CAPITAL LETTER GAMMA
		'\xa7':	'\u0394',	 // GREEK CAPITAL LETTER DELTA
		'\xa8':	'\u0395',	 // GREEK CAPITAL LETTER EPSILON
		'\xa9':	'\u0396',	 // GREEK CAPITAL LETTER ZETA
		'\xaa':	'\u0397',	 // GREEK CAPITAL LETTER ETA
		'\xab':	'\u00bd',	 // VULGAR FRACTION ONE HALF
		'\xac':	'\u0398',	 // GREEK CAPITAL LETTER THETA
		'\xad':	'\u0399',	 // GREEK CAPITAL LETTER IOTA
		'\xae':	'\u00ab',	 // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xaf':	'\u00bb',	 // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xb0':	'\u2591',	 // LIGHT SHADE
		'\xb1':	'\u2592',	 // MEDIUM SHADE
		'\xb2':	'\u2593',	 // DARK SHADE
		'\xb3':	'\u2502',	 // BOX DRAWINGS LIGHT VERTICAL
		'\xb4':	'\u2524',	 // BOX DRAWINGS LIGHT VERTICAL AND LEFT
		'\xb5':	'\u039a',	 // GREEK CAPITAL LETTER KAPPA
		'\xb6':	'\u039b',	 // GREEK CAPITAL LETTER LAMDA
		'\xb7':	'\u039c',	 // GREEK CAPITAL LETTER MU
		'\xb8':	'\u039d',	 // GREEK CAPITAL LETTER NU
		'\xb9':	'\u2563',	 // BOX DRAWINGS DOUBLE VERTICAL AND LEFT
		'\xba':	'\u2551',	 // BOX DRAWINGS DOUBLE VERTICAL
		'\xbb':	'\u2557',	 // BOX DRAWINGS DOUBLE DOWN AND LEFT
		'\xbc':	'\u255d',	 // BOX DRAWINGS DOUBLE UP AND LEFT
		'\xbd':	'\u039e',	 // GREEK CAPITAL LETTER XI
		'\xbe':	'\u039f',	 // GREEK CAPITAL LETTER OMICRON
		'\xbf':	'\u2510',	 // BOX DRAWINGS LIGHT DOWN AND LEFT
		'\xc0':	'\u2514',	 // BOX DRAWINGS LIGHT UP AND RIGHT
		'\xc1':	'\u2534',	 // BOX DRAWINGS LIGHT UP AND HORIZONTAL
		'\xc2':	'\u252c',	 // BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
		'\xc3':	'\u251c',	 // BOX DRAWINGS LIGHT VERTICAL AND RIGHT
		'\xc4':	'\u2500',	 // BOX DRAWINGS LIGHT HORIZONTAL
		'\xc5':	'\u253c',	 // BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
		'\xc6':	'\u03a0',	 // GREEK CAPITAL LETTER PI
		'\xc7':	'\u03a1',	 // GREEK CAPITAL LETTER RHO
		'\xc8':	'\u255a',	 // BOX DRAWINGS DOUBLE UP AND RIGHT
		'\xc9':	'\u2554',	 // BOX DRAWINGS DOUBLE DOWN AND RIGHT
		'\xca':	'\u2569',	 // BOX DRAWINGS DOUBLE UP AND HORIZONTAL
		'\xcb':	'\u2566',	 // BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL
		'\xcc':	'\u2560',	 // BOX DRAWINGS DOUBLE VERTICAL AND RIGHT
		'\xcd':	'\u2550',	 // BOX DRAWINGS DOUBLE HORIZONTAL
		'\xce':	'\u256c',	 // BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL
		'\xcf':	'\u03a3',	 // GREEK CAPITAL LETTER SIGMA
		'\xd0':	'\u03a4',	 // GREEK CAPITAL LETTER TAU
		'\xd1':	'\u03a5',	 // GREEK CAPITAL LETTER UPSILON
		'\xd2':	'\u03a6',	 // GREEK CAPITAL LETTER PHI
		'\xd3':	'\u03a7',	 // GREEK CAPITAL LETTER CHI
		'\xd4':	'\u03a8',	 // GREEK CAPITAL LETTER PSI
		'\xd5':	'\u03a9',	 // GREEK CAPITAL LETTER OMEGA
		'\xd6':	'\u03b1',	 // GREEK SMALL LETTER ALPHA
		'\xd7':	'\u03b2',	 // GREEK SMALL LETTER BETA
		'\xd8':	'\u03b3',	 // GREEK SMALL LETTER GAMMA
		'\xd9':	'\u2518',	 // BOX DRAWINGS LIGHT UP AND LEFT
		'\xda':	'\u250c',	 // BOX DRAWINGS LIGHT DOWN AND RIGHT
		'\xdb':	'\u2588',	 // FULL BLOCK
		'\xdc':	'\u2584',	 // LOWER HALF BLOCK
		'\xdd':	'\u03b4',	 // GREEK SMALL LETTER DELTA
		'\xde':	'\u03b5',	 // GREEK SMALL LETTER EPSILON
		'\xdf':	'\u2580',	 // UPPER HALF BLOCK
		'\xe0':	'\u03b6',	 // GREEK SMALL LETTER ZETA
		'\xe1':	'\u03b7',	 // GREEK SMALL LETTER ETA
		'\xe2':	'\u03b8',	 // GREEK SMALL LETTER THETA
		'\xe3':	'\u03b9',	 // GREEK SMALL LETTER IOTA
		'\xe4':	'\u03ba',	 // GREEK SMALL LETTER KAPPA
		'\xe5':	'\u03bb',	 // GREEK SMALL LETTER LAMDA
		'\xe6':	'\u03bc',	 // GREEK SMALL LETTER MU
		'\xe7':	'\u03bd',	 // GREEK SMALL LETTER NU
		'\xe8':	'\u03be',	 // GREEK SMALL LETTER XI
		'\xe9':	'\u03bf',	 // GREEK SMALL LETTER OMICRON
		'\xea':	'\u03c0',	 // GREEK SMALL LETTER PI
		'\xeb':	'\u03c1',	 // GREEK SMALL LETTER RHO
		'\xec':	'\u03c3',	 // GREEK SMALL LETTER SIGMA
		'\xed':	'\u03c2',	 // GREEK SMALL LETTER FINAL SIGMA
		'\xee':	'\u03c4',	 // GREEK SMALL LETTER TAU
		'\xef':	'\u0384',	 // GREEK TONOS
		'\xf0':	'\u00ad',	 // SOFT HYPHEN
		'\xf1':	'\u00b1',	 // PLUS-MINUS SIGN
		'\xf2':	'\u03c5',	 // GREEK SMALL LETTER UPSILON
		'\xf3':	'\u03c6',	 // GREEK SMALL LETTER PHI
		'\xf4':	'\u03c7',	 // GREEK SMALL LETTER CHI
		'\xf5':	'\u00a7',	 // SECTION SIGN
		'\xf6':	'\u03c8',	 // GREEK SMALL LETTER PSI
		'\xf7':	'\u0385',	 // GREEK DIALYTIKA TONOS
		'\xf8':	'\u00b0',	 // DEGREE SIGN
		'\xf9':	'\u00a8',	 // DIAERESIS
		'\xfa':	'\u03c9',	 // GREEK SMALL LETTER OMEGA
		'\xfb':	'\u03cb',	 // GREEK SMALL LETTER UPSILON WITH DIALYTIKA
		'\xfc':	'\u03b0',	 // GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND TONOS
		'\xfd':	'\u03ce',	 // GREEK SMALL LETTER OMEGA WITH TONOS
		'\xfe':	'\u25a0',	 // BLACK SQUARE
		'\xff':	'\u00a0',	 // NO-BREAK SPACE

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecCP869{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "cp869",
		Aliases: []string{
			"869",
		},
		Codec: codec,
	}

	Register(cm)
}
