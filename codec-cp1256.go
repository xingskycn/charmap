
package charmap

type codecCP1256 struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c codecCP1256) Encode(s string) string {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c codecCP1256) Decode(s string) string {
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
		'\x80':	'\u20AC',	 // EURO SIGN
		'\x81':	'\u067E',	 // ARABIC LETTER PEH
		'\x82':	'\u201A',	 // SINGLE LOW-9 QUOTATION MARK
		'\x83':	'\u0192',	 // LATIN SMALL LETTER F WITH HOOK
		'\x84':	'\u201E',	 // DOUBLE LOW-9 QUOTATION MARK
		'\x85':	'\u2026',	 // HORIZONTAL ELLIPSIS
		'\x86':	'\u2020',	 // DAGGER
		'\x87':	'\u2021',	 // DOUBLE DAGGER
		'\x88':	'\u02C6',	 // MODIFIER LETTER CIRCUMFLEX ACCENT
		'\x89':	'\u2030',	 // PER MILLE SIGN
		'\x8A':	'\u0679',	 // ARABIC LETTER TTEH
		'\x8B':	'\u2039',	 // SINGLE LEFT-POINTING ANGLE QUOTATION MARK
		'\x8C':	'\u0152',	 // LATIN CAPITAL LIGATURE OE
		'\x8D':	'\u0686',	 // ARABIC LETTER TCHEH
		'\x8E':	'\u0698',	 // ARABIC LETTER JEH
		'\x8F':	'\u0688',	 // ARABIC LETTER DDAL
		'\x90':	'\u06AF',	 // ARABIC LETTER GAF
		'\x91':	'\u2018',	 // LEFT SINGLE QUOTATION MARK
		'\x92':	'\u2019',	 // RIGHT SINGLE QUOTATION MARK
		'\x93':	'\u201C',	 // LEFT DOUBLE QUOTATION MARK
		'\x94':	'\u201D',	 // RIGHT DOUBLE QUOTATION MARK
		'\x95':	'\u2022',	 // BULLET
		'\x96':	'\u2013',	 // EN DASH
		'\x97':	'\u2014',	 // EM DASH
		'\x98':	'\u06A9',	 // ARABIC LETTER KEHEH
		'\x99':	'\u2122',	 // TRADE MARK SIGN
		'\x9A':	'\u0691',	 // ARABIC LETTER RREH
		'\x9B':	'\u203A',	 // SINGLE RIGHT-POINTING ANGLE QUOTATION MARK
		'\x9C':	'\u0153',	 // LATIN SMALL LIGATURE OE
		'\x9D':	'\u200C',	 // ZERO WIDTH NON-JOINER
		'\x9E':	'\u200D',	 // ZERO WIDTH JOINER
		'\x9F':	'\u06BA',	 // ARABIC LETTER NOON GHUNNA
		'\xA0':	'\u00A0',	 // NO-BREAK SPACE
		'\xA1':	'\u060C',	 // ARABIC COMMA
		'\xA2':	'\u00A2',	 // CENT SIGN
		'\xA3':	'\u00A3',	 // POUND SIGN
		'\xA4':	'\u00A4',	 // CURRENCY SIGN
		'\xA5':	'\u00A5',	 // YEN SIGN
		'\xA6':	'\u00A6',	 // BROKEN BAR
		'\xA7':	'\u00A7',	 // SECTION SIGN
		'\xA8':	'\u00A8',	 // DIAERESIS
		'\xA9':	'\u00A9',	 // COPYRIGHT SIGN
		'\xAA':	'\u06BE',	 // ARABIC LETTER HEH DOACHASHMEE
		'\xAB':	'\u00AB',	 // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xAC':	'\u00AC',	 // NOT SIGN
		'\xAD':	'\u00AD',	 // SOFT HYPHEN
		'\xAE':	'\u00AE',	 // REGISTERED SIGN
		'\xAF':	'\u00AF',	 // MACRON
		'\xB0':	'\u00B0',	 // DEGREE SIGN
		'\xB1':	'\u00B1',	 // PLUS-MINUS SIGN
		'\xB2':	'\u00B2',	 // SUPERSCRIPT TWO
		'\xB3':	'\u00B3',	 // SUPERSCRIPT THREE
		'\xB4':	'\u00B4',	 // ACUTE ACCENT
		'\xB5':	'\u00B5',	 // MICRO SIGN
		'\xB6':	'\u00B6',	 // PILCROW SIGN
		'\xB7':	'\u00B7',	 // MIDDLE DOT
		'\xB8':	'\u00B8',	 // CEDILLA
		'\xB9':	'\u00B9',	 // SUPERSCRIPT ONE
		'\xBA':	'\u061B',	 // ARABIC SEMICOLON
		'\xBB':	'\u00BB',	 // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xBC':	'\u00BC',	 // VULGAR FRACTION ONE QUARTER
		'\xBD':	'\u00BD',	 // VULGAR FRACTION ONE HALF
		'\xBE':	'\u00BE',	 // VULGAR FRACTION THREE QUARTERS
		'\xBF':	'\u061F',	 // ARABIC QUESTION MARK
		'\xC0':	'\u06C1',	 // ARABIC LETTER HEH GOAL
		'\xC1':	'\u0621',	 // ARABIC LETTER HAMZA
		'\xC2':	'\u0622',	 // ARABIC LETTER ALEF WITH MADDA ABOVE
		'\xC3':	'\u0623',	 // ARABIC LETTER ALEF WITH HAMZA ABOVE
		'\xC4':	'\u0624',	 // ARABIC LETTER WAW WITH HAMZA ABOVE
		'\xC5':	'\u0625',	 // ARABIC LETTER ALEF WITH HAMZA BELOW
		'\xC6':	'\u0626',	 // ARABIC LETTER YEH WITH HAMZA ABOVE
		'\xC7':	'\u0627',	 // ARABIC LETTER ALEF
		'\xC8':	'\u0628',	 // ARABIC LETTER BEH
		'\xC9':	'\u0629',	 // ARABIC LETTER TEH MARBUTA
		'\xCA':	'\u062A',	 // ARABIC LETTER TEH
		'\xCB':	'\u062B',	 // ARABIC LETTER THEH
		'\xCC':	'\u062C',	 // ARABIC LETTER JEEM
		'\xCD':	'\u062D',	 // ARABIC LETTER HAH
		'\xCE':	'\u062E',	 // ARABIC LETTER KHAH
		'\xCF':	'\u062F',	 // ARABIC LETTER DAL
		'\xD0':	'\u0630',	 // ARABIC LETTER THAL
		'\xD1':	'\u0631',	 // ARABIC LETTER REH
		'\xD2':	'\u0632',	 // ARABIC LETTER ZAIN
		'\xD3':	'\u0633',	 // ARABIC LETTER SEEN
		'\xD4':	'\u0634',	 // ARABIC LETTER SHEEN
		'\xD5':	'\u0635',	 // ARABIC LETTER SAD
		'\xD6':	'\u0636',	 // ARABIC LETTER DAD
		'\xD7':	'\u00D7',	 // MULTIPLICATION SIGN
		'\xD8':	'\u0637',	 // ARABIC LETTER TAH
		'\xD9':	'\u0638',	 // ARABIC LETTER ZAH
		'\xDA':	'\u0639',	 // ARABIC LETTER AIN
		'\xDB':	'\u063A',	 // ARABIC LETTER GHAIN
		'\xDC':	'\u0640',	 // ARABIC TATWEEL
		'\xDD':	'\u0641',	 // ARABIC LETTER FEH
		'\xDE':	'\u0642',	 // ARABIC LETTER QAF
		'\xDF':	'\u0643',	 // ARABIC LETTER KAF
		'\xE0':	'\u00E0',	 // LATIN SMALL LETTER A WITH GRAVE
		'\xE1':	'\u0644',	 // ARABIC LETTER LAM
		'\xE2':	'\u00E2',	 // LATIN SMALL LETTER A WITH CIRCUMFLEX
		'\xE3':	'\u0645',	 // ARABIC LETTER MEEM
		'\xE4':	'\u0646',	 // ARABIC LETTER NOON
		'\xE5':	'\u0647',	 // ARABIC LETTER HEH
		'\xE6':	'\u0648',	 // ARABIC LETTER WAW
		'\xE7':	'\u00E7',	 // LATIN SMALL LETTER C WITH CEDILLA
		'\xE8':	'\u00E8',	 // LATIN SMALL LETTER E WITH GRAVE
		'\xE9':	'\u00E9',	 // LATIN SMALL LETTER E WITH ACUTE
		'\xEA':	'\u00EA',	 // LATIN SMALL LETTER E WITH CIRCUMFLEX
		'\xEB':	'\u00EB',	 // LATIN SMALL LETTER E WITH DIAERESIS
		'\xEC':	'\u0649',	 // ARABIC LETTER ALEF MAKSURA
		'\xED':	'\u064A',	 // ARABIC LETTER YEH
		'\xEE':	'\u00EE',	 // LATIN SMALL LETTER I WITH CIRCUMFLEX
		'\xEF':	'\u00EF',	 // LATIN SMALL LETTER I WITH DIAERESIS
		'\xF0':	'\u064B',	 // ARABIC FATHATAN
		'\xF1':	'\u064C',	 // ARABIC DAMMATAN
		'\xF2':	'\u064D',	 // ARABIC KASRATAN
		'\xF3':	'\u064E',	 // ARABIC FATHA
		'\xF4':	'\u00F4',	 // LATIN SMALL LETTER O WITH CIRCUMFLEX
		'\xF5':	'\u064F',	 // ARABIC DAMMA
		'\xF6':	'\u0650',	 // ARABIC KASRA
		'\xF7':	'\u00F7',	 // DIVISION SIGN
		'\xF8':	'\u0651',	 // ARABIC SHADDA
		'\xF9':	'\u00F9',	 // LATIN SMALL LETTER U WITH GRAVE
		'\xFA':	'\u0652',	 // ARABIC SUKUN
		'\xFB':	'\u00FB',	 // LATIN SMALL LETTER U WITH CIRCUMFLEX
		'\xFC':	'\u00FC',	 // LATIN SMALL LETTER U WITH DIAERESIS
		'\xFD':	'\u200E',	 // LEFT-TO-RIGHT MARK
		'\xFE':	'\u200F',	 // RIGHT-TO-LEFT MARK
		'\xFF':	'\u06D2',	 // ARABIC LETTER YEH BARREE

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := codecCP1256{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := Charmap{
		Name: "cp1256",
		Aliases: []string{
			"1256",
		},
		Codec: codec,
	}

	Register(cm)
}
