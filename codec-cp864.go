
package charmap

type codecCP864 struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecCP864) Encode(s string) string {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecCP864) Decode(s string) string {
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
		'\x25':	'\u066a',	 // ARABIC PERCENT SIGN
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
		'\x80':	'\u00b0',	 // DEGREE SIGN
		'\x81':	'\u00b7',	 // MIDDLE DOT
		'\x82':	'\u2219',	 // BULLET OPERATOR
		'\x83':	'\u221a',	 // SQUARE ROOT
		'\x84':	'\u2592',	 // MEDIUM SHADE
		'\x85':	'\u2500',	 // FORMS LIGHT HORIZONTAL
		'\x86':	'\u2502',	 // FORMS LIGHT VERTICAL
		'\x87':	'\u253c',	 // FORMS LIGHT VERTICAL AND HORIZONTAL
		'\x88':	'\u2524',	 // FORMS LIGHT VERTICAL AND LEFT
		'\x89':	'\u252c',	 // FORMS LIGHT DOWN AND HORIZONTAL
		'\x8a':	'\u251c',	 // FORMS LIGHT VERTICAL AND RIGHT
		'\x8b':	'\u2534',	 // FORMS LIGHT UP AND HORIZONTAL
		'\x8c':	'\u2510',	 // FORMS LIGHT DOWN AND LEFT
		'\x8d':	'\u250c',	 // FORMS LIGHT DOWN AND RIGHT
		'\x8e':	'\u2514',	 // FORMS LIGHT UP AND RIGHT
		'\x8f':	'\u2518',	 // FORMS LIGHT UP AND LEFT
		'\x90':	'\u03b2',	 // GREEK SMALL BETA
		'\x91':	'\u221e',	 // INFINITY
		'\x92':	'\u03c6',	 // GREEK SMALL PHI
		'\x93':	'\u00b1',	 // PLUS-OR-MINUS SIGN
		'\x94':	'\u00bd',	 // FRACTION 1/2
		'\x95':	'\u00bc',	 // FRACTION 1/4
		'\x96':	'\u2248',	 // ALMOST EQUAL TO
		'\x97':	'\u00ab',	 // LEFT POINTING GUILLEMET
		'\x98':	'\u00bb',	 // RIGHT POINTING GUILLEMET
		'\x99':	'\ufef7',	 // ARABIC LIGATURE LAM WITH ALEF WITH HAMZA ABOVE ISOLATED FORM
		'\x9a':	'\ufef8',	 // ARABIC LIGATURE LAM WITH ALEF WITH HAMZA ABOVE FINAL FORM
		// '\x9b' UNDEFINED
		// '\x9c' UNDEFINED
		'\x9d':	'\ufefb',	 // ARABIC LIGATURE LAM WITH ALEF ISOLATED FORM
		'\x9e':	'\ufefc',	 // ARABIC LIGATURE LAM WITH ALEF FINAL FORM
		// '\x9f' UNDEFINED
		'\xa0':	'\u00a0',	 // NON-BREAKING SPACE
		'\xa1':	'\u00ad',	 // SOFT HYPHEN
		'\xa2':	'\ufe82',	 // ARABIC LETTER ALEF WITH MADDA ABOVE FINAL FORM
		'\xa3':	'\u00a3',	 // POUND SIGN
		'\xa4':	'\u00a4',	 // CURRENCY SIGN
		'\xa5':	'\ufe84',	 // ARABIC LETTER ALEF WITH HAMZA ABOVE FINAL FORM
		// '\xa6' UNDEFINED
		// '\xa7' UNDEFINED
		'\xa8':	'\ufe8e',	 // ARABIC LETTER ALEF FINAL FORM
		'\xa9':	'\ufe8f',	 // ARABIC LETTER BEH ISOLATED FORM
		'\xaa':	'\ufe95',	 // ARABIC LETTER TEH ISOLATED FORM
		'\xab':	'\ufe99',	 // ARABIC LETTER THEH ISOLATED FORM
		'\xac':	'\u060c',	 // ARABIC COMMA
		'\xad':	'\ufe9d',	 // ARABIC LETTER JEEM ISOLATED FORM
		'\xae':	'\ufea1',	 // ARABIC LETTER HAH ISOLATED FORM
		'\xaf':	'\ufea5',	 // ARABIC LETTER KHAH ISOLATED FORM
		'\xb0':	'\u0660',	 // ARABIC-INDIC DIGIT ZERO
		'\xb1':	'\u0661',	 // ARABIC-INDIC DIGIT ONE
		'\xb2':	'\u0662',	 // ARABIC-INDIC DIGIT TWO
		'\xb3':	'\u0663',	 // ARABIC-INDIC DIGIT THREE
		'\xb4':	'\u0664',	 // ARABIC-INDIC DIGIT FOUR
		'\xb5':	'\u0665',	 // ARABIC-INDIC DIGIT FIVE
		'\xb6':	'\u0666',	 // ARABIC-INDIC DIGIT SIX
		'\xb7':	'\u0667',	 // ARABIC-INDIC DIGIT SEVEN
		'\xb8':	'\u0668',	 // ARABIC-INDIC DIGIT EIGHT
		'\xb9':	'\u0669',	 // ARABIC-INDIC DIGIT NINE
		'\xba':	'\ufed1',	 // ARABIC LETTER FEH ISOLATED FORM
		'\xbb':	'\u061b',	 // ARABIC SEMICOLON
		'\xbc':	'\ufeb1',	 // ARABIC LETTER SEEN ISOLATED FORM
		'\xbd':	'\ufeb5',	 // ARABIC LETTER SHEEN ISOLATED FORM
		'\xbe':	'\ufeb9',	 // ARABIC LETTER SAD ISOLATED FORM
		'\xbf':	'\u061f',	 // ARABIC QUESTION MARK
		'\xc0':	'\u00a2',	 // CENT SIGN
		'\xc1':	'\ufe80',	 // ARABIC LETTER HAMZA ISOLATED FORM
		'\xc2':	'\ufe81',	 // ARABIC LETTER ALEF WITH MADDA ABOVE ISOLATED FORM
		'\xc3':	'\ufe83',	 // ARABIC LETTER ALEF WITH HAMZA ABOVE ISOLATED FORM
		'\xc4':	'\ufe85',	 // ARABIC LETTER WAW WITH HAMZA ABOVE ISOLATED FORM
		'\xc5':	'\ufeca',	 // ARABIC LETTER AIN FINAL FORM
		'\xc6':	'\ufe8b',	 // ARABIC LETTER YEH WITH HAMZA ABOVE INITIAL FORM
		'\xc7':	'\ufe8d',	 // ARABIC LETTER ALEF ISOLATED FORM
		'\xc8':	'\ufe91',	 // ARABIC LETTER BEH INITIAL FORM
		'\xc9':	'\ufe93',	 // ARABIC LETTER TEH MARBUTA ISOLATED FORM
		'\xca':	'\ufe97',	 // ARABIC LETTER TEH INITIAL FORM
		'\xcb':	'\ufe9b',	 // ARABIC LETTER THEH INITIAL FORM
		'\xcc':	'\ufe9f',	 // ARABIC LETTER JEEM INITIAL FORM
		'\xcd':	'\ufea3',	 // ARABIC LETTER HAH INITIAL FORM
		'\xce':	'\ufea7',	 // ARABIC LETTER KHAH INITIAL FORM
		'\xcf':	'\ufea9',	 // ARABIC LETTER DAL ISOLATED FORM
		'\xd0':	'\ufeab',	 // ARABIC LETTER THAL ISOLATED FORM
		'\xd1':	'\ufead',	 // ARABIC LETTER REH ISOLATED FORM
		'\xd2':	'\ufeaf',	 // ARABIC LETTER ZAIN ISOLATED FORM
		'\xd3':	'\ufeb3',	 // ARABIC LETTER SEEN INITIAL FORM
		'\xd4':	'\ufeb7',	 // ARABIC LETTER SHEEN INITIAL FORM
		'\xd5':	'\ufebb',	 // ARABIC LETTER SAD INITIAL FORM
		'\xd6':	'\ufebf',	 // ARABIC LETTER DAD INITIAL FORM
		'\xd7':	'\ufec1',	 // ARABIC LETTER TAH ISOLATED FORM
		'\xd8':	'\ufec5',	 // ARABIC LETTER ZAH ISOLATED FORM
		'\xd9':	'\ufecb',	 // ARABIC LETTER AIN INITIAL FORM
		'\xda':	'\ufecf',	 // ARABIC LETTER GHAIN INITIAL FORM
		'\xdb':	'\u00a6',	 // BROKEN VERTICAL BAR
		'\xdc':	'\u00ac',	 // NOT SIGN
		'\xdd':	'\u00f7',	 // DIVISION SIGN
		'\xde':	'\u00d7',	 // MULTIPLICATION SIGN
		'\xdf':	'\ufec9',	 // ARABIC LETTER AIN ISOLATED FORM
		'\xe0':	'\u0640',	 // ARABIC TATWEEL
		'\xe1':	'\ufed3',	 // ARABIC LETTER FEH INITIAL FORM
		'\xe2':	'\ufed7',	 // ARABIC LETTER QAF INITIAL FORM
		'\xe3':	'\ufedb',	 // ARABIC LETTER KAF INITIAL FORM
		'\xe4':	'\ufedf',	 // ARABIC LETTER LAM INITIAL FORM
		'\xe5':	'\ufee3',	 // ARABIC LETTER MEEM INITIAL FORM
		'\xe6':	'\ufee7',	 // ARABIC LETTER NOON INITIAL FORM
		'\xe7':	'\ufeeb',	 // ARABIC LETTER HEH INITIAL FORM
		'\xe8':	'\ufeed',	 // ARABIC LETTER WAW ISOLATED FORM
		'\xe9':	'\ufeef',	 // ARABIC LETTER ALEF MAKSURA ISOLATED FORM
		'\xea':	'\ufef3',	 // ARABIC LETTER YEH INITIAL FORM
		'\xeb':	'\ufebd',	 // ARABIC LETTER DAD ISOLATED FORM
		'\xec':	'\ufecc',	 // ARABIC LETTER AIN MEDIAL FORM
		'\xed':	'\ufece',	 // ARABIC LETTER GHAIN FINAL FORM
		'\xee':	'\ufecd',	 // ARABIC LETTER GHAIN ISOLATED FORM
		'\xef':	'\ufee1',	 // ARABIC LETTER MEEM ISOLATED FORM
		'\xf0':	'\ufe7d',	 // ARABIC SHADDA MEDIAL FORM
		'\xf1':	'\u0651',	 // ARABIC SHADDAH
		'\xf2':	'\ufee5',	 // ARABIC LETTER NOON ISOLATED FORM
		'\xf3':	'\ufee9',	 // ARABIC LETTER HEH ISOLATED FORM
		'\xf4':	'\ufeec',	 // ARABIC LETTER HEH MEDIAL FORM
		'\xf5':	'\ufef0',	 // ARABIC LETTER ALEF MAKSURA FINAL FORM
		'\xf6':	'\ufef2',	 // ARABIC LETTER YEH FINAL FORM
		'\xf7':	'\ufed0',	 // ARABIC LETTER GHAIN MEDIAL FORM
		'\xf8':	'\ufed5',	 // ARABIC LETTER QAF ISOLATED FORM
		'\xf9':	'\ufef5',	 // ARABIC LIGATURE LAM WITH ALEF WITH MADDA ABOVE ISOLATED FORM
		'\xfa':	'\ufef6',	 // ARABIC LIGATURE LAM WITH ALEF WITH MADDA ABOVE FINAL FORM
		'\xfb':	'\ufedd',	 // ARABIC LETTER LAM ISOLATED FORM
		'\xfc':	'\ufed9',	 // ARABIC LETTER KAF ISOLATED FORM
		'\xfd':	'\ufef1',	 // ARABIC LETTER YEH ISOLATED FORM
		'\xfe':	'\u25a0',	 // BLACK SQUARE
		// '\xff' UNDEFINED

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecCP864{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "cp864",
		Aliases: []string{
			"864",
		},
		Codec: codec,
	}

	Register(cm)
}
