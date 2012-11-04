
package charmap

type codecMAC_GREEK struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecMAC_GREEK) Encode(s string) (string, error) {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecMAC_GREEK) Decode(s string) (string, error) {
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
		'\x80':	'\u00C4',	 // LATIN CAPITAL LETTER A WITH DIAERESIS
		'\x81':	'\u00B9',	 // SUPERSCRIPT ONE
		'\x82':	'\u00B2',	 // SUPERSCRIPT TWO
		'\x83':	'\u00C9',	 // LATIN CAPITAL LETTER E WITH ACUTE
		'\x84':	'\u00B3',	 // SUPERSCRIPT THREE
		'\x85':	'\u00D6',	 // LATIN CAPITAL LETTER O WITH DIAERESIS
		'\x86':	'\u00DC',	 // LATIN CAPITAL LETTER U WITH DIAERESIS
		'\x87':	'\u0385',	 // GREEK DIALYTIKA TONOS
		'\x88':	'\u00E0',	 // LATIN SMALL LETTER A WITH GRAVE
		'\x89':	'\u00E2',	 // LATIN SMALL LETTER A WITH CIRCUMFLEX
		'\x8A':	'\u00E4',	 // LATIN SMALL LETTER A WITH DIAERESIS
		'\x8B':	'\u0384',	 // GREEK TONOS
		'\x8C':	'\u00A8',	 // DIAERESIS
		'\x8D':	'\u00E7',	 // LATIN SMALL LETTER C WITH CEDILLA
		'\x8E':	'\u00E9',	 // LATIN SMALL LETTER E WITH ACUTE
		'\x8F':	'\u00E8',	 // LATIN SMALL LETTER E WITH GRAVE
		'\x90':	'\u00EA',	 // LATIN SMALL LETTER E WITH CIRCUMFLEX
		'\x91':	'\u00EB',	 // LATIN SMALL LETTER E WITH DIAERESIS
		'\x92':	'\u00A3',	 // POUND SIGN
		'\x93':	'\u2122',	 // TRADE MARK SIGN
		'\x94':	'\u00EE',	 // LATIN SMALL LETTER I WITH CIRCUMFLEX
		'\x95':	'\u00EF',	 // LATIN SMALL LETTER I WITH DIAERESIS
		'\x96':	'\u2022',	 // BULLET
		'\x97':	'\u00BD',	 // VULGAR FRACTION ONE HALF
		'\x98':	'\u2030',	 // PER MILLE SIGN
		'\x99':	'\u00F4',	 // LATIN SMALL LETTER O WITH CIRCUMFLEX
		'\x9A':	'\u00F6',	 // LATIN SMALL LETTER O WITH DIAERESIS
		'\x9B':	'\u00A6',	 // BROKEN BAR
		'\x9C':	'\u00AD',	 // SOFT HYPHEN
		'\x9D':	'\u00F9',	 // LATIN SMALL LETTER U WITH GRAVE
		'\x9E':	'\u00FB',	 // LATIN SMALL LETTER U WITH CIRCUMFLEX
		'\x9F':	'\u00FC',	 // LATIN SMALL LETTER U WITH DIAERESIS
		'\xA0':	'\u2020',	 // DAGGER
		'\xA1':	'\u0393',	 // GREEK CAPITAL LETTER GAMMA
		'\xA2':	'\u0394',	 // GREEK CAPITAL LETTER DELTA
		'\xA3':	'\u0398',	 // GREEK CAPITAL LETTER THETA
		'\xA4':	'\u039B',	 // GREEK CAPITAL LETTER LAMBDA
		'\xA5':	'\u039E',	 // GREEK CAPITAL LETTER XI
		'\xA6':	'\u03A0',	 // GREEK CAPITAL LETTER PI
		'\xA7':	'\u00DF',	 // LATIN SMALL LETTER SHARP S
		'\xA8':	'\u00AE',	 // REGISTERED SIGN
		'\xA9':	'\u00A9',	 // COPYRIGHT SIGN
		'\xAA':	'\u03A3',	 // GREEK CAPITAL LETTER SIGMA
		'\xAB':	'\u03AA',	 // GREEK CAPITAL LETTER IOTA WITH DIALYTIKA
		'\xAC':	'\u00A7',	 // SECTION SIGN
		'\xAD':	'\u2260',	 // NOT EQUAL TO
		'\xAE':	'\u00B0',	 // DEGREE SIGN
		'\xAF':	'\u0387',	 // GREEK ANO TELEIA
		'\xB0':	'\u0391',	 // GREEK CAPITAL LETTER ALPHA
		'\xB1':	'\u00B1',	 // PLUS-MINUS SIGN
		'\xB2':	'\u2264',	 // LESS-THAN OR EQUAL TO
		'\xB3':	'\u2265',	 // GREATER-THAN OR EQUAL TO
		'\xB4':	'\u00A5',	 // YEN SIGN
		'\xB5':	'\u0392',	 // GREEK CAPITAL LETTER BETA
		'\xB6':	'\u0395',	 // GREEK CAPITAL LETTER EPSILON
		'\xB7':	'\u0396',	 // GREEK CAPITAL LETTER ZETA
		'\xB8':	'\u0397',	 // GREEK CAPITAL LETTER ETA
		'\xB9':	'\u0399',	 // GREEK CAPITAL LETTER IOTA
		'\xBA':	'\u039A',	 // GREEK CAPITAL LETTER KAPPA
		'\xBB':	'\u039C',	 // GREEK CAPITAL LETTER MU
		'\xBC':	'\u03A6',	 // GREEK CAPITAL LETTER PHI
		'\xBD':	'\u03AB',	 // GREEK CAPITAL LETTER UPSILON WITH DIALYTIKA
		'\xBE':	'\u03A8',	 // GREEK CAPITAL LETTER PSI
		'\xBF':	'\u03A9',	 // GREEK CAPITAL LETTER OMEGA
		'\xC0':	'\u03AC',	 // GREEK SMALL LETTER ALPHA WITH TONOS
		'\xC1':	'\u039D',	 // GREEK CAPITAL LETTER NU
		'\xC2':	'\u00AC',	 // NOT SIGN
		'\xC3':	'\u039F',	 // GREEK CAPITAL LETTER OMICRON
		'\xC4':	'\u03A1',	 // GREEK CAPITAL LETTER RHO
		'\xC5':	'\u2248',	 // ALMOST EQUAL TO
		'\xC6':	'\u03A4',	 // GREEK CAPITAL LETTER TAU
		'\xC7':	'\u00AB',	 // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xC8':	'\u00BB',	 // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xC9':	'\u2026',	 // HORIZONTAL ELLIPSIS
		'\xCA':	'\u00A0',	 // NO-BREAK SPACE
		'\xCB':	'\u03A5',	 // GREEK CAPITAL LETTER UPSILON
		'\xCC':	'\u03A7',	 // GREEK CAPITAL LETTER CHI
		'\xCD':	'\u0386',	 // GREEK CAPITAL LETTER ALPHA WITH TONOS
		'\xCE':	'\u0388',	 // GREEK CAPITAL LETTER EPSILON WITH TONOS
		'\xCF':	'\u0153',	 // LATIN SMALL LIGATURE OE
		'\xD0':	'\u2013',	 // EN DASH
		'\xD1':	'\u2015',	 // HORIZONTAL BAR
		'\xD2':	'\u201C',	 // LEFT DOUBLE QUOTATION MARK
		'\xD3':	'\u201D',	 // RIGHT DOUBLE QUOTATION MARK
		'\xD4':	'\u2018',	 // LEFT SINGLE QUOTATION MARK
		'\xD5':	'\u2019',	 // RIGHT SINGLE QUOTATION MARK
		'\xD6':	'\u00F7',	 // DIVISION SIGN
		'\xD7':	'\u0389',	 // GREEK CAPITAL LETTER ETA WITH TONOS
		'\xD8':	'\u038A',	 // GREEK CAPITAL LETTER IOTA WITH TONOS
		'\xD9':	'\u038C',	 // GREEK CAPITAL LETTER OMICRON WITH TONOS
		'\xDA':	'\u038E',	 // GREEK CAPITAL LETTER UPSILON WITH TONOS
		'\xDB':	'\u03AD',	 // GREEK SMALL LETTER EPSILON WITH TONOS
		'\xDC':	'\u03AE',	 // GREEK SMALL LETTER ETA WITH TONOS
		'\xDD':	'\u03AF',	 // GREEK SMALL LETTER IOTA WITH TONOS
		'\xDE':	'\u03CC',	 // GREEK SMALL LETTER OMICRON WITH TONOS
		'\xDF':	'\u038F',	 // GREEK CAPITAL LETTER OMEGA WITH TONOS
		'\xE0':	'\u03CD',	 // GREEK SMALL LETTER UPSILON WITH TONOS
		'\xE1':	'\u03B1',	 // GREEK SMALL LETTER ALPHA
		'\xE2':	'\u03B2',	 // GREEK SMALL LETTER BETA
		'\xE3':	'\u03C8',	 // GREEK SMALL LETTER PSI
		'\xE4':	'\u03B4',	 // GREEK SMALL LETTER DELTA
		'\xE5':	'\u03B5',	 // GREEK SMALL LETTER EPSILON
		'\xE6':	'\u03C6',	 // GREEK SMALL LETTER PHI
		'\xE7':	'\u03B3',	 // GREEK SMALL LETTER GAMMA
		'\xE8':	'\u03B7',	 // GREEK SMALL LETTER ETA
		'\xE9':	'\u03B9',	 // GREEK SMALL LETTER IOTA
		'\xEA':	'\u03BE',	 // GREEK SMALL LETTER XI
		'\xEB':	'\u03BA',	 // GREEK SMALL LETTER KAPPA
		'\xEC':	'\u03BB',	 // GREEK SMALL LETTER LAMBDA
		'\xED':	'\u03BC',	 // GREEK SMALL LETTER MU
		'\xEE':	'\u03BD',	 // GREEK SMALL LETTER NU
		'\xEF':	'\u03BF',	 // GREEK SMALL LETTER OMICRON
		'\xF0':	'\u03C0',	 // GREEK SMALL LETTER PI
		'\xF1':	'\u03CE',	 // GREEK SMALL LETTER OMEGA WITH TONOS
		'\xF2':	'\u03C1',	 // GREEK SMALL LETTER RHO
		'\xF3':	'\u03C3',	 // GREEK SMALL LETTER SIGMA
		'\xF4':	'\u03C4',	 // GREEK SMALL LETTER TAU
		'\xF5':	'\u03B8',	 // GREEK SMALL LETTER THETA
		'\xF6':	'\u03C9',	 // GREEK SMALL LETTER OMEGA
		'\xF7':	'\u03C2',	 // GREEK SMALL LETTER FINAL SIGMA
		'\xF8':	'\u03C7',	 // GREEK SMALL LETTER CHI
		'\xF9':	'\u03C5',	 // GREEK SMALL LETTER UPSILON
		'\xFA':	'\u03B6',	 // GREEK SMALL LETTER ZETA
		'\xFB':	'\u03CA',	 // GREEK SMALL LETTER IOTA WITH DIALYTIKA
		'\xFC':	'\u03CB',	 // GREEK SMALL LETTER UPSILON WITH DIALYTIKA
		'\xFD':	'\u0390',	 // GREEK SMALL LETTER IOTA WITH DIALYTIKA AND TONOS
		'\xFE':	'\u03B0',	 // GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND TONOS
		// '\xFF' UNDEFINED

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecMAC_GREEK{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "mac_greek",
		Aliases: []string{
			"macgreek",
		},
		Codec: codec,
	}

	Register(cm)
}
