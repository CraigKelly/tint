package main

import (
	"regexp"
	"unicode"
	"unicode/utf8"
)

// TODO: special module just for cursing, lgbtq

// BadTerm represents a single failed existence check and supports matching
// (using a FileMap) via the Match function.
type BadTerm struct {
	SearchTerm string
	Message    string
}

// helper to check for unicode letter at given index
func indexIsWordChar(s string, idx int) bool {
	if idx < 0 || idx >= len(s) {
		return false
	}
	r, _ := utf8.DecodeRuneInString(s[idx:])
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// Match checks test for SearchTerm and returns all matches found
func (bt *BadTerm) Match(fm *FileMap) []*Warning {
	rx := regexp.MustCompile(bt.SearchTerm)
	found := rx.FindAllStringIndex(fm.CheckText, -1)
	if found == nil {
		return nil
	}

	ret := make([]*Warning, 0, len(found))
	for _, loc := range found {
		// make sure we are at word boundaries
		if indexIsWordChar(fm.Text, loc[0]-1) || indexIsWordChar(fm.Text, loc[1]) {
			continue
		}
		// found something
		ret = append(ret, NewWarning(fm, loc[0], bt.Message, fm.Text[loc[0]:loc[1]]))
	}

	return ret
}

// ShouldNotExist returns a slice of BadTerm's, none of which should be in the
// a text (case insensitive). The BadTerm message expects the finder to format
// with the matching string portion (not the SearchTerm).
func ShouldNotExist() []TextCheck {
	// For now this is just hard-coded (until we have ways of enabling or
	// disabling certain checks)
	return []TextCheck{
		// Skunked terms
		&BadTerm{"bona fides", "'%s' has issues - rephrase."},
		&BadTerm{"deceptively", "'%s' has issues - rephrase."},
		&BadTerm{"decimate", "'%s' has issues - rephrase."},
		&BadTerm{"effete", "'%s' has issues - rephrase."},
		&BadTerm{"fulsome", "'%s' has issues - rephrase."},
		&BadTerm{"hopefully", "'%s' has issues - rephrase."},
		&BadTerm{"impassionate", "'%s' has issues - rephrase."},
		&BadTerm{"thankfully,", "'%s' has issues - rephrase."},

		// Airlinese
		&BadTerm{"enplan(ed|es|e|ing|ement)", "'%s' is airlinese."},
		&BadTerm{"deplan(ed|es|e|ing|ement)", "'%s' is airlinese."},

		// Archaic
		&BadTerm{"alack", "'%s' is archaic."},
		&BadTerm{"anent", "'%s' is archaic."},
		&BadTerm{"begat", "'%s' is archaic."},
		&BadTerm{"belike", "'%s' is archaic."},
		&BadTerm{"betimes", "'%s' is archaic."},
		&BadTerm{"boughten", "'%s' is archaic."},
		&BadTerm{"brocage", "'%s' is archaic."},
		&BadTerm{"brokage", "'%s' is archaic."},
		&BadTerm{"camarade", "'%s' is archaic."},
		&BadTerm{"chiefer", "'%s' is archaic."},
		&BadTerm{"chiefest", "'%s' is archaic."},
		&BadTerm{"christiana", "'%s' is archaic."},
		&BadTerm{"completely obsolescent", "'%s' is archaic."},
		&BadTerm{"cozen", "'%s' is archaic."},
		&BadTerm{"divers", "'%s' is archaic."},
		&BadTerm{"deflexion", "'%s' is archaic."},
		&BadTerm{"durst", "'%s' is archaic."},
		&BadTerm{"fain", "'%s' is archaic."},
		&BadTerm{"forsooth", "'%s' is archaic."},
		&BadTerm{"foreclose from", "'%s' is archaic."},
		&BadTerm{"haply", "'%s' is archaic."},
		&BadTerm{"howbeit", "'%s' is archaic."},
		&BadTerm{"illumine", "'%s' is archaic."},
		&BadTerm{"in sooth", "'%s' is archaic."},
		&BadTerm{"maugre", "'%s' is archaic."},
		&BadTerm{"meseems", "'%s' is archaic."},
		&BadTerm{"nigh", "'%s' is archaic."},
		&BadTerm{"peradventure", "'%s' is archaic."},
		&BadTerm{"perchance", "'%s' is archaic."},
		&BadTerm{"saith", "'%s' is archaic."},
		&BadTerm{"shew", "'%s' is archaic."},
		&BadTerm{"sistren", "'%s' is archaic."},
		&BadTerm{"spake", "'%s' is archaic."},
		&BadTerm{"to wit", "'%s' is archaic."},
		&BadTerm{"verily", "'%s' is archaic."},
		&BadTerm{"whilom", "'%s' is archaic."},
		&BadTerm{"withal", "'%s' is archaic."},
		&BadTerm{"wot", "'%s' is archaic."},
		&BadTerm{"enclosed please find", "'%s' is archaic."},
		&BadTerm{"please find enclosed", "'%s' is archaic."},
		&BadTerm{"enclosed herewith", "'%s' is archaic."},
		&BadTerm{"enclosed herein", "'%s' is archaic."},
		&BadTerm{"inforce", "'%s' is archaic."},
		&BadTerm{"ex postfacto", "'%s' is archaic."},
		&BadTerm{"foreclose from", "'%s' is archaic."},
		&BadTerm{"forewent", "'%s' is archaic."},
		&BadTerm{"for ever", "'%s' is archaic."},

		// Corporate speak
		&BadTerm{"at the end of the day", "Minimize corporatism like '%s'"},
		&BadTerm{"back to the drawing board", "Minimize corporatism like '%s'"},
		&BadTerm{"hit the ground running", "Minimize corporatism like '%s'"},
		&BadTerm{"get the ball rolling", "Minimize corporatism like '%s'"},
		&BadTerm{"low-hanging fruit", "Minimize corporatism like '%s'"},
		&BadTerm{"thrown under the bus", "Minimize corporatism like '%s'"},
		&BadTerm{"think outside the box", "Minimize corporatism like '%s'"},
		&BadTerm{"touch base", "Minimize corporatism like '%s'"},
		&BadTerm{"get my manager's blessing", "Minimize corporatism like '%s'"},
		&BadTerm{"it's on my radar", "Minimize corporatism like '%s'"},
		&BadTerm{"ping me", "Minimize corporatism like '%s'"},
		&BadTerm{"i don't have the bandwidth", "Minimize corporatism like '%s'"},
		&BadTerm{"no brainer", "Minimize corporatism like '%s'"},
		&BadTerm{"par for the course", "Minimize corporatism like '%s'"},
		&BadTerm{"bang for (the|your) buck", "Minimize corporatism like '%s'"},
		&BadTerm{"synerg(y|istic)", "Minimize corporatism like '%s'"},
		&BadTerm{"apples to apples", "Minimize corporatism like '%s'"},
		&BadTerm{"win-win", "Minimize corporatism like '%s'"},
		&BadTerm{"circle back around", "Minimize corporatism like '%s'"},
		&BadTerm{"all hands on deck", "Minimize corporatism like '%s'"},
		&BadTerm{"drill-down", "Minimize corporatism like '%s'"},
		&BadTerm{"elephant in the room", "Minimize corporatism like '%s'"},
		&BadTerm{"meets? with your approval", "'%s' is bureaucratese."},

		// Date/Time
		&BadTerm{`12 ?[ap]\.?m\.?`, "'%s' can be confusing, use midnight or noon"},
		&BadTerm{`\d{1,2} ?a\.?m\.? in the morning`, "'%s' is redundant - AM is always morning"},
		&BadTerm{`\d{1,2} ?p\.?m\.? in the evening`, "'%s' is redundant - PM is always night"},
		&BadTerm{`\d{1,2} ?p\.?m\.? at night`, "'%s' is redundant - PM is always night"},
		&BadTerm{`\d{1,2} ?p\.?m\.? in the afternoon`, "'%s' is redundant - PM is always after noon"},

		// Hedging
		&BadTerm{"i would argue that", "Don't hedge with '%s', just say it"},
		&BadTerm{"so to speak", "Don't hedge with '%s', just say it"},
		&BadTerm{"to a certain degree", "Don't hedge with '%s', just say it"},

		// Jargon
		&BadTerm{"in the affirmative", "'%s' is jargon, say it another way."},
		&BadTerm{"in the negative", "'%s' is jargon, say it another way."},
		&BadTerm{"agendize", "'%s' is jargon, say it another way."},
		&BadTerm{"per your order", "'%s' is jargon, say it another way."},
		&BadTerm{"per your request", "'%s' is jargon, say it another way."},
		&BadTerm{"disincentivize", "'%s' is jargon, say it another way."},
		&BadTerm{"reconceptualize", "'%s' is jargon and the hallmark of a pretentious ass."},
		&BadTerm{"demassification", "'%s' is jargon and the hallmark of a pretentious ass."},
		&BadTerm{"attitudinally", "'%s' is jargon and the hallmark of a pretentious ass unless you are making a specific kind of geometricpoint."},
		&BadTerm{"judgementally", "'%s' is jargon and the hallmark of a pretentious ass."},

		// Lexical illusion
		&BadTerm{`the\sthe`, "Repeated words - '%s' is probably a lexical illusion"},
		&BadTerm{`am\sam`, "Repeated words - '%s' is probably a lexical illusion"},
		&BadTerm{`has\shas`, "Repeated words - '%s' is probably a lexical illusion"},

		// Malapropisms
		&BadTerm{"the infinitesimal universe", "'%s' is a malapropism"},
		&BadTerm{"a serial experienc", "'%s' is a malapropism"},
		&BadTerm{"attack my voracity", "'%s' is a malapropism"},

		// Mixed metaphor: bottle necks
		&BadTerm{"biggest bottleneck", "'%s' is a mixed metaphor - bottle with big necks are easy to pass through"},
		&BadTerm{"big bottleneck", "'%s' is a mixed metaphor - bottle with big necks are easy to pass through"},
		&BadTerm{"large bottleneck", "'%s' is a mixed metaphor - bottle with big necks are easy to pass through"},
		&BadTerm{"largest bottleneck", "'%s' is a mixed metaphor - bottle with big necks are easy to pass through"},
		&BadTerm{"world-wide bottleneck", "'%s' is a mixed metaphor - bottle with big necks are easy to pass through"},
		&BadTerm{"huge bottleneck", "'%s' is a mixed metaphor - bottle with big necks are easy to pass through"},
		&BadTerm{"massive bottleneck", "'%s' is a mixed metaphor - bottle with big necks are easy to pass through"},

		// Oxymorons
		&BadTerm{"amateur expert", "'%s' is an oxymoron"},
		&BadTerm{"increasingly less", "'%s' is an oxymoron"},
		&BadTerm{"advancing backwards?", "'%s' is an oxymoron"},
		&BadTerm{"alludes explicitly to", "'%s' is an oxymoron"},
		&BadTerm{"explicitly alludes to", "'%s' is an oxymoron"},
		&BadTerm{"totally obsolescent", "'%s' is an oxymoron"},
		&BadTerm{"completely obsolescent", "'%s' is an oxymoron"},
		&BadTerm{"generally always", "'%s' is an oxymoron"},
		&BadTerm{"usually always", "'%s' is an oxymoron"},
		&BadTerm{"increasingly less", "'%s' is an oxymoron"},
		&BadTerm{"build down", "'%s' is an oxymoron"},
		&BadTerm{"conspicuous absence", "'%s' is an oxymoron"},
		&BadTerm{"exact estimate", "'%s' is an oxymoron"},
		&BadTerm{"found missing", "'%s' is an oxymoron"},
		&BadTerm{"intense apathy", "'%s' is an oxymoron"},
		&BadTerm{"mandatory choice", "'%s' is an oxymoron"},
		&BadTerm{"nonworking mother", "'%s' is an oxymoron"},
		&BadTerm{"organized mess", "'%s' is an oxymoron"},

		// Psychology
		&BadTerm{`p ?\= ?0\.0+`, "Unless '%s' really means zero you should use more decimal places."},

		// Security - credit card numbers don't need to be in text. Patterns from proselint
		// and the IIN section of https://en.wikipedia.org/wiki/Payment_card_number
		// VISA: start with 4 and have length of 13, 16, 19
		&BadTerm{`4\d{12}`, "'%s' appears to be a VISA credit card number. Seriously?"},
		&BadTerm{`4\d{15}`, "'%s' appears to be a VISA credit card number. Seriously?"},
		&BadTerm{`4\d{18}`, "'%s' appears to be a VISA credit card number. Seriously?"},
		// MasterCard: 2221-2720 or 51-55 with 16 digits
		&BadTerm{`2[2-7]\d{14}`, "'%s' appears to be a MasterCard credit card number. Seriously?"},
		&BadTerm{`5[1-5]\d{14}`, "'%s' appears to be a MasterCard credit card number. Seriously?"},
		// AmEx 34,37 with 15 digits
		&BadTerm{`3[4,7]\d{13}`, "'%s' appears to be an AmEx credit card number. Seriously?"},
		&BadTerm{`3[0,6,8]\d{12}`, "'%s' appears to be an AmEx credit card number. Seriously?"},
		// Discover: 6011, 622126-622925, 644-649, 65 with len of 16 or 19
		&BadTerm{`6011\d{12}`, "'%s' appears to be a Discover credit card number. Seriously?"},
		&BadTerm{`6011\d{15}`, "'%s' appears to be a Discover credit card number. Seriously?"},
		&BadTerm{`622\d{13}`, "'%s' appears to be a Discover credit card number. Seriously?"},
		&BadTerm{`622\d{16}`, "'%s' appears to be a Discover credit card number. Seriously?"},
		&BadTerm{`64[4-9]\d{13}`, "'%s' appears to be a Discover credit card number. Seriously?"},
		&BadTerm{`64[4-9]\d{16}`, "'%s' appears to be a Discover credit card number. Seriously?"},
		&BadTerm{`65\d{14}`, "'%s' appears to be a Discover credit card number. Seriously?"},
		&BadTerm{`65\d{17}`, "'%s' appears to be a Discover credit card number. Seriously?"},
		// JCB: 3528-3589 with len of 16
		&BadTerm{`35[2-8]\d{13}`, "'%s' appears to be a JCB credit card number. Seriously?"},
		// Dankort: 5019, 4175, 4571 with len of 16 (note some caught by VISA)
		&BadTerm{`5019\d{12}`, "'%s' appears to be a Dankort credit card number. Seriously?"},

		// Sexism
		&BadTerm{"anchorman", "'%s' shows gender bias."},
		&BadTerm{"anchorwoman", "'%s' shows gender bias."},
		&BadTerm{"chairman", "'%s' shows gender bias."},
		&BadTerm{"chairwoman", "'%s' shows gender bias."},
		&BadTerm{"draftman", "'%s' shows gender bias."},
		&BadTerm{"draftwoman", "'%s' shows gender bias."},
		&BadTerm{"ombudsman", "'%s' shows gender bias."},
		&BadTerm{"ombudswoman", "'%s' shows gender bias."},
		&BadTerm{"tribesman", "'%s' shows gender bias."},
		&BadTerm{"tribeswoman", "'%s' shows gender bias."},
		&BadTerm{"policeman", "'%s' shows gender bias."},
		&BadTerm{"policewoman", "'%s' shows gender bias."},
		&BadTerm{"fireman", "'%s' shows gender bias."},
		&BadTerm{"firewoman", "'%s' shows gender bias."},
		&BadTerm{"mailman", "'%s' shows gender bias."},
		&BadTerm{"mailwoman", "'%s' shows gender bias."},
		&BadTerm{"herstory", "'%s' shows gender bias."},
		&BadTerm{"womyn", "'%s' shows gender bias."},
		&BadTerm{"poetess", "'%s' shows gender bias."},
		&BadTerm{"authoress", "'%s' shows gender bias."},
		&BadTerm{"waitress", "'%s' shows gender bias."},
		&BadTerm{"lady lawyer", "'%s' shows gender bias."},
		&BadTerm{"woman doctor", "'%s' shows gender bias."},
		&BadTerm{"female booksalesman", "'%s' shows gender bias."},
		&BadTerm{"female airman", "'%s' shows gender bias."},
		&BadTerm{"executrix", "'%s' shows gender bias."},
		&BadTerm{"prosecutrix", "'%s' shows gender bias."},
		&BadTerm{"testatrix", "'%s' shows gender bias."},
		&BadTerm{"man and wife", "'%s' shows gender bias."},
		&BadTerm{"chairmen and chairs", "'%s' shows gender bias."},
		&BadTerm{"men and girls", "'%s' shows gender bias."},
		&BadTerm{"comedienne", "'%s' shows gender bias."},
		&BadTerm{"confidante", "'%s' shows gender bias."},
		&BadTerm{"woman scientist", "'%s' shows gender bias."},
		&BadTerm{"women scientists", "'%s' shows gender bias."},
		&BadTerm{"heroine", "'%s' shows gender bias."},

		// Weasel words
		&BadTerm{"very", "'%s' is a weasel word"},
		&BadTerm{"extremely", "'%s' is a weasel word"},
		&BadTerm{"various", "'%s' is a weasel word"},
		&BadTerm{"a number of", "'%s' is a weasel word"},
		&BadTerm{"fairly", "'%s' is a weasel word"},
		&BadTerm{"quite", "'%s' is a weasel word"},
		&BadTerm{"interestingly", "'%s' is a weasel word"},
		&BadTerm{"surprisingly", "'%s' is a weasel word"},
		&BadTerm{"remarkably", "'%s' is a weasel word"},
		&BadTerm{"clearly", "'%s' is a weasel word"},

		// Not words
		&BadTerm{"doubtlessly", "'%s' is not a word."},
		&BadTerm{"analyzation", "'%s' is not a word."},
		&BadTerm{"annoyment", "'%s' is not a word."},
		&BadTerm{"confirmant", "'%s' is not a word."},
		&BadTerm{"confirmants", "'%s' is not a word."},
		&BadTerm{"conversate", "'%s' is not a word."},
		&BadTerm{"crained", "'%s' is not a word."},
		&BadTerm{"dispersement", "'%s' is not a word."},
		&BadTerm{"discomforture", "'%s' is not a word."},
		&BadTerm{"affrontery", "'%s' is not a word."},
		&BadTerm{"forebearance", "'%s' is not a word."},
		&BadTerm{"improprietous", "'%s' is not a word."},
		&BadTerm{"inclimate", "'%s' is not a word."},
		&BadTerm{"relative inexpense", "'%s' is not a word."},
		&BadTerm{"inimicable", "'%s' is not a word."},
		&BadTerm{"irregardless", "'%s' is not a word."},
		&BadTerm{"minimalize", "'%s' is not a word."},
		&BadTerm{"minimalized", "'%s' is not a word."},
		&BadTerm{"minimalizes", "'%s' is not a word."},
		&BadTerm{"minimalizing", "'%s' is not a word."},
		&BadTerm{"optimalize", "'%s' is not a word."},
		&BadTerm{"paralyzation", "'%s' is not a word."},
		&BadTerm{"pettifogger", "'%s' is not a word."},
		&BadTerm{"proprietous", "'%s' is not a word."},
		&BadTerm{"squelch", "'%s' is not a word."},
		&BadTerm{"seldomly", "'%s' is not a word."},
		&BadTerm{"thusly", "'%s' is not a word."},
		&BadTerm{"uncategorically", "'%s' is not a word."},
		&BadTerm{"undoubtably", "'%s' is not a word."},
		&BadTerm{"unequivocable", "'%s' is not a word."},
		&BadTerm{"unmercilessly", "'%s' is not a word."},
		&BadTerm{"unrelentlessly", "'%s' is not a word."},

		// Bad usage
		&BadTerm{`et\. al\.?`, "'%s' has misplaced punctuation. It is 'et al'"},
		&BadTerm{"suddenly,", "'%s' slows the action and warns your reader."},
		&BadTerm{"a not unjustifiable assumption", "'%s' is considered bad usage."},
		&BadTerm{"leaves much to be desired", "'%s' is considered bad usage."},
		&BadTerm{"would serve no purpose", "'%s' is considered bad usage."},
		&BadTerm{"a consideration which we should do well to bear in mind", "'%s' is considered bad usage."},
		&BadTerm{"obviously", "Do not use '%s'"},
		&BadTerm{"utilize", "Do not use '%s'"},
		&BadTerm{"inferior to", "'%s' is not a true comparative. Use 'to' instead of 'than'"},
		&BadTerm{"superior to", "'%s' is not a true comparative. Use 'to' instead of 'than'"},
		&BadTerm{"not guilty beyond (a |any )?reasonable doubt", "Avoid '%s' unless you are a lawyer and a copy-editor. Almost always a miscue."},
		&BadTerm{"between you and i", "Don't use '%s'."},
		&BadTerm{"on accident", "Don't use '%s'."},
		&BadTerm{"somewhat of a", "Don't use '%s'."},
		&BadTerm{"all it'?s own", "Don't use '%s'."},
		&BadTerm{"reason is because", "Don't use '%s'."},
		&BadTerm{"in regards to", "Don't use '%s'."},
		&BadTerm{"would of", "Don't use '%s'."},
		&BadTerm{"i ?(?:feel|am feeling|am|'m|'m feeling) nauseous", "Don't use '%s'. Also reconsider using nauseated."},
		&BadTerm{"from whence", "You don't need 'from' in '%s'."},

		// False plurals
		&BadTerm{"many kudos", "'%s' is incorrect - kudos is singular."},
		&BadTerm{"talismen", "'%s' is an incorrect plural"},
		&BadTerm{"phenomenons", "'%s' is an incorrect plural"},

		// Illogics
		&BadTerm{"preplan", "'%s' makes no sense"},
		&BadTerm{"appraisal valuations?", "'%s' makes no sense"},
		&BadTerm{"(?:i|you|he|she|it|y'all|all y'all|you all|they) could care less", "'%s' makes no sense"},
		&BadTerm{"least worst", "'%s' makes no sense"},
		&BadTerm{"much-needed gaps?", "'%s' makes no sense"},
		&BadTerm{"much-needed voids?", "'%s' makes no sense"},
		&BadTerm{"no longer requires oxygen", "'%s' makes no sense"},
		&BadTerm{"without scarcely", "'%s' makes no sense"},
		&BadTerm{"to coin a phrase from", "'%s' makes no sense. Did you mean 'borrow'?"},
		&BadTerm{"without your collusion", "'%s' makes no sense. Did you you mean 'acquiescence'?"},

		// Metadiscourse
		&BadTerm{"the preceeding discussion", "'%s' is metadiscourse. Avoid it."},
		&BadTerm{"the rest of this article", "'%s' is metadiscourse. Avoid it."},
		&BadTerm{"this chapter discusses", "'%s' is metadiscourse. Avoid it."},
		&BadTerm{"the preceding paragraph demonstrated", "'%s' is metadiscourse. Avoid it."},
		&BadTerm{"the previous section analyzed", "'%s' is metadiscourse. Avoid it."},

		// Waxing X requires X=adjective, not adverb (so you wax enthusiastic)
		&BadTerm{"wax(es|ed|ing)? [[:alpha:]]+ly", "'%s' is incorrect - use an adjective instead of an adverb."},
	}
}
