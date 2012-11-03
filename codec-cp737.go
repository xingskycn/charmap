
package charmap

type codecCP737 struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecCP737) Encode(s string) string {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecCP737) Decode(s string) string {
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
		'\x80':	'\u0391',	 // GREEK CAPITAL LETTER ALPHA
		'\x81':	'\u0392',	 // GREEK CAPITAL LETTER BETA
		'\x82':	'\u0393',	 // GREEK CAPITAL LETTER GAMMA
		'\x83':	'\u0394',	 // GREEK CAPITAL LETTER DELTA
		'\x84':	'\u0395',	 // GREEK CAPITAL LETTER EPSILON
		'\x85':	'\u0396',	 // GREEK CAPITAL LETTER ZETA
		'\x86':	'\u0397',	 // GREEK CAPITAL LETTER ETA
		'\x87':	'\u0398',	 // GREEK CAPITAL LETTER THETA
		'\x88':	'\u0399',	 // GREEK CAPITAL LETTER IOTA
		'\x89':	'\u039a',	 // GREEK CAPITAL LETTER KAPPA
		'\x8a':	'\u039b',	 // GREEK CAPITAL LETTER LAMDA
		'\x8b':	'\u039c',	 // GREEK CAPITAL LETTER MU
		'\x8c':	'\u039d',	 // GREEK CAPITAL LETTER NU
		'\x8d':	'\u039e',	 // GREEK CAPITAL LETTER XI
		'\x8e':	'\u039f',	 // GREEK CAPITAL LETTER OMICRON
		'\x8f':	'\u03a0',	 // GREEK CAPITAL LETTER PI
		'\x90':	'\u03a1',	 // GREEK CAPITAL LETTER RHO
		'\x91':	'\u03a3',	 // GREEK CAPITAL LETTER SIGMA
		'\x92':	'\u03a4',	 // GREEK CAPITAL LETTER TAU
		'\x93':	'\u03a5',	 // GREEK CAPITAL LETTER UPSILON
		'\x94':	'\u03a6',	 // GREEK CAPITAL LETTER PHI
		'\x95':	'\u03a7',	 // GREEK CAPITAL LETTER CHI
		'\x96':	'\u03a8',	 // GREEK CAPITAL LETTER PSI
		'\x97':	'\u03a9',	 // GREEK CAPITAL LETTER OMEGA
		'\x98':	'\u03b1',	 // GREEK SMALL LETTER ALPHA
		'\x99':	'\u03b2',	 // GREEK SMALL LETTER BETA
		'\x9a':	'\u03b3',	 // GREEK SMALL LETTER GAMMA
		'\x9b':	'\u03b4',	 // GREEK SMALL LETTER DELTA
		'\x9c':	'\u03b5',	 // GREEK SMALL LETTER EPSILON
		'\x9d':	'\u03b6',	 // GREEK SMALL LETTER ZETA
		'\x9e':	'\u03b7',	 // GREEK SMALL LETTER ETA
		'\x9f':	'\u03b8',	 // GREEK SMALL LETTER THETA
		'\xa0':	'\u03b9',	 // GREEK SMALL LETTER IOTA
		'\xa1':	'\u03ba',	 // GREEK SMALL LETTER KAPPA
		'\xa2':	'\u03bb',	 // GREEK SMALL LETTER LAMDA
		'\xa3':	'\u03bc',	 // GREEK SMALL LETTER MU
		'\xa4':	'\u03bd',	 // GREEK SMALL LETTER NU
		'\xa5':	'\u03be',	 // GREEK SMALL LETTER XI
		'\xa6':	'\u03bf',	 // GREEK SMALL LETTER OMICRON
		'\xa7':	'\u03c0',	 // GREEK SMALL LETTER PI
		'\xa8':	'\u03c1',	 // GREEK SMALL LETTER RHO
		'\xa9':	'\u03c3',	 // GREEK SMALL LETTER SIGMA
		'\xaa':	'\u03c2',	 // GREEK SMALL LETTER FINAL SIGMA
		'\xab':	'\u03c4',	 // GREEK SMALL LETTER TAU
		'\xac':	'\u03c5',	 // GREEK SMALL LETTER UPSILON
		'\xad':	'\u03c6',	 // GREEK SMALL LETTER PHI
		'\xae':	'\u03c7',	 // GREEK SMALL LETTER CHI
		'\xaf':	'\u03c8',	 // GREEK SMALL LETTER PSI
		'\xb0':	'\u2591',	 // LIGHT SHADE
		'\xb1':	'\u2592',	 // MEDIUM SHADE
		'\xb2':	'\u2593',	 // DARK SHADE
		'\xb3':	'\u2502',	 // BOX DRAWINGS LIGHT VERTICAL
		'\xb4':	'\u2524',	 // BOX DRAWINGS LIGHT VERTICAL AND LEFT
		'\xb5':	'\u2561',	 // BOX DRAWINGS VERTICAL SINGLE AND LEFT DOUBLE
		'\xb6':	'\u2562',	 // BOX DRAWINGS VERTICAL DOUBLE AND LEFT SINGLE
		'\xb7':	'\u2556',	 // BOX DRAWINGS DOWN DOUBLE AND LEFT SINGLE
		'\xb8':	'\u2555',	 // BOX DRAWINGS DOWN SINGLE AND LEFT DOUBLE
		'\xb9':	'\u2563',	 // BOX DRAWINGS DOUBLE VERTICAL AND LEFT
		'\xba':	'\u2551',	 // BOX DRAWINGS DOUBLE VERTICAL
		'\xbb':	'\u2557',	 // BOX DRAWINGS DOUBLE DOWN AND LEFT
		'\xbc':	'\u255d',	 // BOX DRAWINGS DOUBLE UP AND LEFT
		'\xbd':	'\u255c',	 // BOX DRAWINGS UP DOUBLE AND LEFT SINGLE
		'\xbe':	'\u255b',	 // BOX DRAWINGS UP SINGLE AND LEFT DOUBLE
		'\xbf':	'\u2510',	 // BOX DRAWINGS LIGHT DOWN AND LEFT
		'\xc0':	'\u2514',	 // BOX DRAWINGS LIGHT UP AND RIGHT
		'\xc1':	'\u2534',	 // BOX DRAWINGS LIGHT UP AND HORIZONTAL
		'\xc2':	'\u252c',	 // BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
		'\xc3':	'\u251c',	 // BOX DRAWINGS LIGHT VERTICAL AND RIGHT
		'\xc4':	'\u2500',	 // BOX DRAWINGS LIGHT HORIZONTAL
		'\xc5':	'\u253c',	 // BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
		'\xc6':	'\u255e',	 // BOX DRAWINGS VERTICAL SINGLE AND RIGHT DOUBLE
		'\xc7':	'\u255f',	 // BOX DRAWINGS VERTICAL DOUBLE AND RIGHT SINGLE
		'\xc8':	'\u255a',	 // BOX DRAWINGS DOUBLE UP AND RIGHT
		'\xc9':	'\u2554',	 // BOX DRAWINGS DOUBLE DOWN AND RIGHT
		'\xca':	'\u2569',	 // BOX DRAWINGS DOUBLE UP AND HORIZONTAL
		'\xcb':	'\u2566',	 // BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL
		'\xcc':	'\u2560',	 // BOX DRAWINGS DOUBLE VERTICAL AND RIGHT
		'\xcd':	'\u2550',	 // BOX DRAWINGS DOUBLE HORIZONTAL
		'\xce':	'\u256c',	 // BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL
		'\xcf':	'\u2567',	 // BOX DRAWINGS UP SINGLE AND HORIZONTAL DOUBLE
		'\xd0':	'\u2568',	 // BOX DRAWINGS UP DOUBLE AND HORIZONTAL SINGLE
		'\xd1':	'\u2564',	 // BOX DRAWINGS DOWN SINGLE AND HORIZONTAL DOUBLE
		'\xd2':	'\u2565',	 // BOX DRAWINGS DOWN DOUBLE AND HORIZONTAL SINGLE
		'\xd3':	'\u2559',	 // BOX DRAWINGS UP DOUBLE AND RIGHT SINGLE
		'\xd4':	'\u2558',	 // BOX DRAWINGS UP SINGLE AND RIGHT DOUBLE
		'\xd5':	'\u2552',	 // BOX DRAWINGS DOWN SINGLE AND RIGHT DOUBLE
		'\xd6':	'\u2553',	 // BOX DRAWINGS DOWN DOUBLE AND RIGHT SINGLE
		'\xd7':	'\u256b',	 // BOX DRAWINGS VERTICAL DOUBLE AND HORIZONTAL SINGLE
		'\xd8':	'\u256a',	 // BOX DRAWINGS VERTICAL SINGLE AND HORIZONTAL DOUBLE
		'\xd9':	'\u2518',	 // BOX DRAWINGS LIGHT UP AND LEFT
		'\xda':	'\u250c',	 // BOX DRAWINGS LIGHT DOWN AND RIGHT
		'\xdb':	'\u2588',	 // FULL BLOCK
		'\xdc':	'\u2584',	 // LOWER HALF BLOCK
		'\xdd':	'\u258c',	 // LEFT HALF BLOCK
		'\xde':	'\u2590',	 // RIGHT HALF BLOCK
		'\xdf':	'\u2580',	 // UPPER HALF BLOCK
		'\xe0':	'\u03c9',	 // GREEK SMALL LETTER OMEGA
		'\xe1':	'\u03ac',	 // GREEK SMALL LETTER ALPHA WITH TONOS
		'\xe2':	'\u03ad',	 // GREEK SMALL LETTER EPSILON WITH TONOS
		'\xe3':	'\u03ae',	 // GREEK SMALL LETTER ETA WITH TONOS
		'\xe4':	'\u03ca',	 // GREEK SMALL LETTER IOTA WITH DIALYTIKA
		'\xe5':	'\u03af',	 // GREEK SMALL LETTER IOTA WITH TONOS
		'\xe6':	'\u03cc',	 // GREEK SMALL LETTER OMICRON WITH TONOS
		'\xe7':	'\u03cd',	 // GREEK SMALL LETTER UPSILON WITH TONOS
		'\xe8':	'\u03cb',	 // GREEK SMALL LETTER UPSILON WITH DIALYTIKA
		'\xe9':	'\u03ce',	 // GREEK SMALL LETTER OMEGA WITH TONOS
		'\xea':	'\u0386',	 // GREEK CAPITAL LETTER ALPHA WITH TONOS
		'\xeb':	'\u0388',	 // GREEK CAPITAL LETTER EPSILON WITH TONOS
		'\xec':	'\u0389',	 // GREEK CAPITAL LETTER ETA WITH TONOS
		'\xed':	'\u038a',	 // GREEK CAPITAL LETTER IOTA WITH TONOS
		'\xee':	'\u038c',	 // GREEK CAPITAL LETTER OMICRON WITH TONOS
		'\xef':	'\u038e',	 // GREEK CAPITAL LETTER UPSILON WITH TONOS
		'\xf0':	'\u038f',	 // GREEK CAPITAL LETTER OMEGA WITH TONOS
		'\xf1':	'\u00b1',	 // PLUS-MINUS SIGN
		'\xf2':	'\u2265',	 // GREATER-THAN OR EQUAL TO
		'\xf3':	'\u2264',	 // LESS-THAN OR EQUAL TO
		'\xf4':	'\u03aa',	 // GREEK CAPITAL LETTER IOTA WITH DIALYTIKA
		'\xf5':	'\u03ab',	 // GREEK CAPITAL LETTER UPSILON WITH DIALYTIKA
		'\xf6':	'\u00f7',	 // DIVISION SIGN
		'\xf7':	'\u2248',	 // ALMOST EQUAL TO
		'\xf8':	'\u00b0',	 // DEGREE SIGN
		'\xf9':	'\u2219',	 // BULLET OPERATOR
		'\xfa':	'\u00b7',	 // MIDDLE DOT
		'\xfb':	'\u221a',	 // SQUARE ROOT
		'\xfc':	'\u207f',	 // SUPERSCRIPT LATIN SMALL LETTER N
		'\xfd':	'\u00b2',	 // SUPERSCRIPT TWO
		'\xfe':	'\u25a0',	 // BLACK SQUARE
		'\xff':	'\u00a0',	 // NO-BREAK SPACE

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecCP737{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "cp737",
		Aliases: []string{
			"737",
		},
		Codec: codec,
	}

	Register(cm)
}