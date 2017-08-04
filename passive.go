package main

import (
	"regexp"
	"strings"
)

// NEEDCAPS

type PassiveTerm struct {
	*BadTerm
}

func (pt *PassiveTerm) Match(fm *FileMap) []*Warning {
	rx := regexp.MustCompile(pt.SearchTerm)
	found := rx.FindAllStringIndex(fm.CheckText, -1)
	if found == nil {
		return nil
	}

	ret := make([]*Warning, 0)
	for _, loc := range found {
		// must be at word bound
		if indexIsWordChar(fm.CheckText, loc[0]-1) || indexIsWordChar(fm.CheckText, loc[1]) {
			continue
		}

		// now find previous word
		prev_end := loc[0] - 2
		for prev_end >= 0 && !indexIsWordChar(fm.CheckText, prev_end) {
			prev_end -= 1
		}
		prev_start := prev_end - 1
		for prev_start >= 0 && indexIsWordChar(fm.CheckText, prev_start) {
			prev_start -= 1
		}
		if prev_start < 0 || prev_end < 0 {
			continue
		}
		prev_word := strings.TrimSpace(fm.CheckText[prev_start : prev_end+1])
		if len(prev_word) < 2 {
			continue
		}

		// Is the previous word one we are looking for?
		prev_bad := false
		verbs := []string{"am", "are", "were", "being", "is", "been", "was", "be"}
		for _, v := range verbs {
			if prev_word == v {
				prev_bad = true
				break
			}
		}

		if prev_bad {
			ret = append(ret, NewWarning(fm, prev_start, pt.Message, fm.Text[prev_start:loc[1]]))
		}
	}

	return ret
}

func passive(secondary string) TextCheck {
	return &PassiveTerm{&BadTerm{secondary, "'%s' appears to be in the passive voice."}}
}

// ShouldNotCliche returns a slice of BadTerm's, none of which should be in the
// a text (case insensitive). See existence_checks.go for details of BadTerms.
func ShouldNotPassive() []TextCheck {
	return []TextCheck{
		passive("ate"),
		passive("awoke"),
		passive("awoken"),
		passive("beat"),
		passive("beaten"),
		passive("become"),
		passive("been"),
		passive("began"),
		passive("begun"),
		passive("bent"),
		passive("beset"),
		passive("bet"),
		passive("bid"),
		passive("bidden"),
		passive("bit"),
		passive("bitten"),
		passive("bled"),
		passive("blew"),
		passive("blown"),
		passive("born"),
		passive("bought"),
		passive("bound"),
		passive("bred"),
		passive("broadcast"),
		passive("broke"),
		passive("broken"),
		passive("brought"),
		passive("built"),
		passive("burnt"),
		passive("burst"),
		passive("cast"),
		passive("caught"),
		passive("chose"),
		passive("chosen"),
		passive("clung"),
		passive("come"),
		passive("cost"),
		passive("crept"),
		passive("cut"),
		passive("dealt"),
		passive("did"),
		passive("dived"),
		passive("done"),
		passive("drawn"),
		passive("dreamt"),
		passive("drew"),
		passive("driven"),
		passive("drove"),
		passive("drunk"),
		passive("dug"),
		passive("eaten"),
		passive("fallen"),
		passive("fed"),
		passive("felt"),
		passive("fit"),
		passive("fled"),
		passive("flown"),
		passive("flung"),
		passive("forbade"),
		passive("forbidden"),
		passive("foregone"),
		passive("forgave"),
		passive("forgiven"),
		passive("forgot"),
		passive("forgotten"),
		passive("forsaken"),
		passive("fought"),
		passive("found"),
		passive("froze"),
		passive("frozen"),
		passive("gave"),
		passive("given"),
		passive("gone"),
		passive("got"),
		passive("gotten"),
		passive("grinded"),
		passive("ground"),
		passive("grown"),
		passive("heard"),
		passive("held"),
		passive("hid"),
		passive("hidden"),
		passive("hit"),
		passive("hung"),
		passive("hurt"),
		passive("kept"),
		passive("knelt"),
		passive("knew"),
		passive("knit"),
		passive("known"),
		passive("laid"),
		passive("lain"),
		passive("leapt"),
		passive("learnt"),
		passive("led"),
		passive("left"),
		passive("lent"),
		passive("let"),
		passive("lighted"),
		passive("lost"),
		passive("made"),
		passive("meant"),
		passive("met"),
		passive("misspelt"),
		passive("mistaken"),
		passive("mown"),
		passive("overcome"),
		passive("overdone"),
		passive("overtaken"),
		passive("overthrown"),
		passive("paid"),
		passive("pled"),
		passive("proved"),
		passive("proven"),
		passive("put"),
		passive("quit"),
		passive("ran"),
		passive("rang"),
		passive("read"),
		passive("rid"),
		passive("ridden"),
		passive("risen"),
		passive("rode"),
		passive("run"),
		passive("rung"),
		passive("said"),
		passive("sat"),
		passive("saw"),
		passive("sawn"),
		passive("seen"),
		passive("sent"),
		passive("set"),
		passive("sewn"),
		passive("shaken"),
		passive("shaved"),
		passive("shaven"),
		passive("shed"),
		passive("shod"),
		passive("shone"),
		passive("shook"),
		passive("shorn"),
		passive("shot"),
		passive("shown"),
		passive("shrunk"),
		passive("shut"),
		passive("slain"),
		passive("slept"),
		passive("slew"),
		passive("slid"),
		passive("slit"),
		passive("slung"),
		passive("smitten"),
		passive("sold"),
		passive("sought"),
		passive("sown"),
		passive("sped"),
		passive("spent"),
		passive("spilt"),
		passive("spit"),
		passive("split"),
		passive("spoke"),
		passive("spoken"),
		passive("spread"),
		passive("sprung"),
		passive("spun"),
		passive("stole"),
		passive("stolen"),
		passive("stood"),
		passive("stridden"),
		passive("striven"),
		passive("struck"),
		passive("strung"),
		passive("stuck"),
		passive("stung"),
		passive("stunk"),
		passive("sung"),
		passive("sunk"),
		passive("swept"),
		passive("swollen"),
		passive("sworn"),
		passive("swum"),
		passive("swung"),
		passive("taken"),
		passive("taught"),
		passive("thought"),
		passive("threw"),
		passive("thrived"),
		passive("thrown"),
		passive("thrust"),
		passive("told"),
		passive("took"),
		passive("tore"),
		passive("torn"),
		passive("trodden"),
		passive("undergone"),
		passive("understood"),
		passive("underwent"),
		passive("upheld"),
		passive("upset"),
		passive("wed"),
		passive("wept"),
		passive("withdrawn"),
		passive("withdrew"),
		passive("withheld"),
		passive("withstood"),
		passive("woke"),
		passive("woken"),
		passive("won"),
		passive("wore"),
		passive("worn"),
		passive("wound"),
		passive("woven"),
		passive("written"),
		passive("wrote"),
		passive("wrung"),
	}
}
