
package charmap

type codecMAC_TURKISH struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecMAC_TURKISH) Encode(s string) string {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecMAC_TURKISH) Decode(s string) string {
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
		'\x81':	'\u00C5',	 // LATIN CAPITAL LETTER A WITH RING ABOVE
		'\x82':	'\u00C7',	 // LATIN CAPITAL LETTER C WITH CEDILLA
		'\x83':	'\u00C9',	 // LATIN CAPITAL LETTER E WITH ACUTE
		'\x84':	'\u00D1',	 // LATIN CAPITAL LETTER N WITH TILDE
		'\x85':	'\u00D6',	 // LATIN CAPITAL LETTER O WITH DIAERESIS
		'\x86':	'\u00DC',	 // LATIN CAPITAL LETTER U WITH DIAERESIS
		'\x87':	'\u00E1',	 // LATIN SMALL LETTER A WITH ACUTE
		'\x88':	'\u00E0',	 // LATIN SMALL LETTER A WITH GRAVE
		'\x89':	'\u00E2',	 // LATIN SMALL LETTER A WITH CIRCUMFLEX
		'\x8A':	'\u00E4',	 // LATIN SMALL LETTER A WITH DIAERESIS
		'\x8B':	'\u00E3',	 // LATIN SMALL LETTER A WITH TILDE
		'\x8C':	'\u00E5',	 // LATIN SMALL LETTER A WITH RING ABOVE
		'\x8D':	'\u00E7',	 // LATIN SMALL LETTER C WITH CEDILLA
		'\x8E':	'\u00E9',	 // LATIN SMALL LETTER E WITH ACUTE
		'\x8F':	'\u00E8',	 // LATIN SMALL LETTER E WITH GRAVE
		'\x90':	'\u00EA',	 // LATIN SMALL LETTER E WITH CIRCUMFLEX
		'\x91':	'\u00EB',	 // LATIN SMALL LETTER E WITH DIAERESIS
		'\x92':	'\u00ED',	 // LATIN SMALL LETTER I WITH ACUTE
		'\x93':	'\u00EC',	 // LATIN SMALL LETTER I WITH GRAVE
		'\x94':	'\u00EE',	 // LATIN SMALL LETTER I WITH CIRCUMFLEX
		'\x95':	'\u00EF',	 // LATIN SMALL LETTER I WITH DIAERESIS
		'\x96':	'\u00F1',	 // LATIN SMALL LETTER N WITH TILDE
		'\x97':	'\u00F3',	 // LATIN SMALL LETTER O WITH ACUTE
		'\x98':	'\u00F2',	 // LATIN SMALL LETTER O WITH GRAVE
		'\x99':	'\u00F4',	 // LATIN SMALL LETTER O WITH CIRCUMFLEX
		'\x9A':	'\u00F6',	 // LATIN SMALL LETTER O WITH DIAERESIS
		'\x9B':	'\u00F5',	 // LATIN SMALL LETTER O WITH TILDE
		'\x9C':	'\u00FA',	 // LATIN SMALL LETTER U WITH ACUTE
		'\x9D':	'\u00F9',	 // LATIN SMALL LETTER U WITH GRAVE
		'\x9E':	'\u00FB',	 // LATIN SMALL LETTER U WITH CIRCUMFLEX
		'\x9F':	'\u00FC',	 // LATIN SMALL LETTER U WITH DIAERESIS
		'\xA0':	'\u2020',	 // DAGGER
		'\xA1':	'\u00B0',	 // DEGREE SIGN
		'\xA2':	'\u00A2',	 // CENT SIGN
		'\xA3':	'\u00A3',	 // POUND SIGN
		'\xA4':	'\u00A7',	 // SECTION SIGN
		'\xA5':	'\u2022',	 // BULLET
		'\xA6':	'\u00B6',	 // PILCROW SIGN
		'\xA7':	'\u00DF',	 // LATIN SMALL LETTER SHARP S
		'\xA8':	'\u00AE',	 // REGISTERED SIGN
		'\xA9':	'\u00A9',	 // COPYRIGHT SIGN
		'\xAA':	'\u2122',	 // TRADE MARK SIGN
		'\xAB':	'\u00B4',	 // ACUTE ACCENT
		'\xAC':	'\u00A8',	 // DIAERESIS
		'\xAD':	'\u2260',	 // NOT EQUAL TO
		'\xAE':	'\u00C6',	 // LATIN CAPITAL LIGATURE AE
		'\xAF':	'\u00D8',	 // LATIN CAPITAL LETTER O WITH STROKE
		'\xB0':	'\u221E',	 // INFINITY
		'\xB1':	'\u00B1',	 // PLUS-MINUS SIGN
		'\xB2':	'\u2264',	 // LESS-THAN OR EQUAL TO
		'\xB3':	'\u2265',	 // GREATER-THAN OR EQUAL TO
		'\xB4':	'\u00A5',	 // YEN SIGN
		'\xB5':	'\u00B5',	 // MICRO SIGN
		'\xB6':	'\u2202',	 // PARTIAL DIFFERENTIAL
		'\xB7':	'\u2211',	 // N-ARY SUMMATION
		'\xB8':	'\u220F',	 // N-ARY PRODUCT
		'\xB9':	'\u03C0',	 // GREEK SMALL LETTER PI
		'\xBA':	'\u222B',	 // INTEGRAL
		'\xBB':	'\u00AA',	 // FEMININE ORDINAL INDICATOR
		'\xBC':	'\u00BA',	 // MASCULINE ORDINAL INDICATOR
		'\xBD':	'\u2126',	 // OHM SIGN
		'\xBE':	'\u00E6',	 // LATIN SMALL LIGATURE AE
		'\xBF':	'\u00F8',	 // LATIN SMALL LETTER O WITH STROKE
		'\xC0':	'\u00BF',	 // INVERTED QUESTION MARK
		'\xC1':	'\u00A1',	 // INVERTED EXCLAMATION MARK
		'\xC2':	'\u00AC',	 // NOT SIGN
		'\xC3':	'\u221A',	 // SQUARE ROOT
		'\xC4':	'\u0192',	 // LATIN SMALL LETTER F WITH HOOK
		'\xC5':	'\u2248',	 // ALMOST EQUAL TO
		'\xC6':	'\u2206',	 // INCREMENT
		'\xC7':	'\u00AB',	 // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xC8':	'\u00BB',	 // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xC9':	'\u2026',	 // HORIZONTAL ELLIPSIS
		'\xCA':	'\u00A0',	 // NO-BREAK SPACE
		'\xCB':	'\u00C0',	 // LATIN CAPITAL LETTER A WITH GRAVE
		'\xCC':	'\u00C3',	 // LATIN CAPITAL LETTER A WITH TILDE
		'\xCD':	'\u00D5',	 // LATIN CAPITAL LETTER O WITH TILDE
		'\xCE':	'\u0152',	 // LATIN CAPITAL LIGATURE OE
		'\xCF':	'\u0153',	 // LATIN SMALL LIGATURE OE
		'\xD0':	'\u2013',	 // EN DASH
		'\xD1':	'\u2014',	 // EM DASH
		'\xD2':	'\u201C',	 // LEFT DOUBLE QUOTATION MARK
		'\xD3':	'\u201D',	 // RIGHT DOUBLE QUOTATION MARK
		'\xD4':	'\u2018',	 // LEFT SINGLE QUOTATION MARK
		'\xD5':	'\u2019',	 // RIGHT SINGLE QUOTATION MARK
		'\xD6':	'\u00F7',	 // DIVISION SIGN
		'\xD7':	'\u25CA',	 // LOZENGE
		'\xD8':	'\u00FF',	 // LATIN SMALL LETTER Y WITH DIAERESIS
		'\xD9':	'\u0178',	 // LATIN CAPITAL LETTER Y WITH DIAERESIS
		'\xDA':	'\u011E',	 // LATIN CAPITAL LETTER G WITH BREVE
		'\xDB':	'\u011F',	 // LATIN SMALL LETTER G WITH BREVE
		'\xDC':	'\u0130',	 // LATIN CAPITAL LETTER I WITH DOT ABOVE
		'\xDD':	'\u0131',	 // LATIN SMALL LETTER DOTLESS I
		'\xDE':	'\u015E',	 // LATIN CAPITAL LETTER S WITH CEDILLA
		'\xDF':	'\u015F',	 // LATIN SMALL LETTER S WITH CEDILLA
		'\xE0':	'\u2021',	 // DOUBLE DAGGER
		'\xE1':	'\u00B7',	 // MIDDLE DOT
		'\xE2':	'\u201A',	 // SINGLE LOW-9 QUOTATION MARK
		'\xE3':	'\u201E',	 // DOUBLE LOW-9 QUOTATION MARK
		'\xE4':	'\u2030',	 // PER MILLE SIGN
		'\xE5':	'\u00C2',	 // LATIN CAPITAL LETTER A WITH CIRCUMFLEX
		'\xE6':	'\u00CA',	 // LATIN CAPITAL LETTER E WITH CIRCUMFLEX
		'\xE7':	'\u00C1',	 // LATIN CAPITAL LETTER A WITH ACUTE
		'\xE8':	'\u00CB',	 // LATIN CAPITAL LETTER E WITH DIAERESIS
		'\xE9':	'\u00C8',	 // LATIN CAPITAL LETTER E WITH GRAVE
		'\xEA':	'\u00CD',	 // LATIN CAPITAL LETTER I WITH ACUTE
		'\xEB':	'\u00CE',	 // LATIN CAPITAL LETTER I WITH CIRCUMFLEX
		'\xEC':	'\u00CF',	 // LATIN CAPITAL LETTER I WITH DIAERESIS
		'\xED':	'\u00CC',	 // LATIN CAPITAL LETTER I WITH GRAVE
		'\xEE':	'\u00D3',	 // LATIN CAPITAL LETTER O WITH ACUTE
		'\xEF':	'\u00D4',	 // LATIN CAPITAL LETTER O WITH CIRCUMFLEX
		// '\xF0' UNDEFINED
		'\xF1':	'\u00D2',	 // LATIN CAPITAL LETTER O WITH GRAVE
		'\xF2':	'\u00DA',	 // LATIN CAPITAL LETTER U WITH ACUTE
		'\xF3':	'\u00DB',	 // LATIN CAPITAL LETTER U WITH CIRCUMFLEX
		'\xF4':	'\u00D9',	 // LATIN CAPITAL LETTER U WITH GRAVE
		// '\xF5' UNDEFINED
		'\xF6':	'\u02C6',	 // MODIFIER LETTER CIRCUMFLEX ACCENT
		'\xF7':	'\u02DC',	 // SMALL TILDE
		'\xF8':	'\u00AF',	 // MACRON
		'\xF9':	'\u02D8',	 // BREVE
		'\xFA':	'\u02D9',	 // DOT ABOVE
		'\xFB':	'\u02DA',	 // RING ABOVE
		'\xFC':	'\u00B8',	 // CEDILLA
		'\xFD':	'\u02DD',	 // DOUBLE ACUTE ACCENT
		'\xFE':	'\u02DB',	 // OGONEK
		'\xFF':	'\u02C7',	 // CARON

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecMAC_TURKISH{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "mac_turkish",
		Aliases: []string{
			"macturkish",
		},
		Codec: codec,
	}

	Register(cm)
}
