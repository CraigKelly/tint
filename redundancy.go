package main

// NEEDCAPS

// ShouldNotRedundant returns a slice of BadTerm's, none of which should be in the
// a text (case insensitive).
func ShouldNotRedundant() []TextCheck {
	// For now this is just hard-coded (until we have ways of enabling or
	// disabling certain checks)
	return []TextCheck{
		&BadTerm{"abm missile", "'%s' is redundant. Just use ABM"},
		&BadTerm{"act test", "'%s' is redundant. Just use ACT"},
		&BadTerm{"abm missiles", "'%s' is redundant. Just use ABM"},
		&BadTerm{"abs braking system", "'%s' is redundant. Just use ABS"},
		&BadTerm{"atm machine", "'%s' is redundant. Just use ATM"},
		&BadTerm{"cd disc", "'%s' is redundant. Just use CD"},
		&BadTerm{"cpi index", "'%s' is redundant. Just use CPI"},
		&BadTerm{"gps system", "'%s' is redundant. Just use GPS"},
		&BadTerm{"gui interface", "'%s' is redundant. Just use GUI"},
		&BadTerm{"hiv virus", "'%s' is redundant. Just use HIV"},
		&BadTerm{"isbn number", "'%s' is redundant. Just use ISBN"},
		&BadTerm{"lcd display", "'%s' is redundant. Just use LCD"},
		&BadTerm{"pdf format", "'%s' is redundant. Just use PDF"},
		&BadTerm{"pin number", "'%s' is redundant. Just use PIN"},
		&BadTerm{"ras syndrome", "'%s' is redundant. Just use RAS"},
		&BadTerm{"rip in peace", "'%s' is redundant. Just use RIP"},
		&BadTerm{"please rsvp", "'%s' is redundant. Just use RSVP"},
		&BadTerm{"salt talks", "'%s' is redundant. Just use SALT"},
		&BadTerm{"sat test", "'%s' is redundant. Just use SAT"},
		&BadTerm{"upc codes", "'%s' is redundant. Just use UPC"},

		&BadTerm{"audible to the ear", "'%s' is redundant. Just use audible"},
		&BadTerm{"rectangular in shape", "'%s' is redundant. Just use rectangular"},

		&BadTerm{"adequate enough", "'%s' is redundant."},
		&BadTerm{"self-admitted", "'%s' is redundant."},
		&BadTerm{"sworn affidavit", "'%s' is redundant."},
		&BadTerm{"mutual agreement", "'%s' is redundant."},
		&BadTerm{"former alumnus", "'%s' is redundant."},
		&BadTerm{"directly antithetical", "'%s' is redundant."},
		&BadTerm{"approximately about", "'%s' is redundant."},
		&BadTerm{"temporary bivouac", "'%s' is redundant."},
		&BadTerm{"bivouac camp", "'%s' is redundant."},
		&BadTerm{"blend together", "'%s' is redundant."},
		&BadTerm{"but nevertheless", "'%s' is redundant."},
		&BadTerm{"accused of a charge", "'%s' is redundant."},
		&BadTerm{"circumstances surrounding", "'%s' is redundant."},
		&BadTerm{"surrounding circumstances", "'%s' is redundant."},
		&BadTerm{"close proximity", "'%s' is redundant."},
		&BadTerm{"collaborate together", "'%s' is redundant."},
		&BadTerm{"fellow collaborator", "'%s' is redundant."},
		&BadTerm{"fellow collaborators", "'%s' is redundant."},
		&BadTerm{"collocated together", "'%s' is redundant."},
		&BadTerm{"fellow colleagues", "'%s' is redundant."},
		&BadTerm{"combine together", "'%s' is redundant."},
		&BadTerm{"self-complacent", "'%s' is redundant."},
		&BadTerm{"self-confessed", "'%s' is redundant."},
		&BadTerm{"connect together", "'%s' is redundant."},
		&BadTerm{"general consensus", "'%s' is redundant."},
		&BadTerm{"consolidate together", "'%s' is redundant."},
		&BadTerm{"still continues to", "'%s' is redundant."},
		&BadTerm{"mutually contradictory", "'%s' is redundant."},
		&BadTerm{"mutual cooperation", "'%s' is redundant."},
		&BadTerm{"couple together", "'%s' is redundant."},
		&BadTerm{"serious crisis", "'%s' is redundant."},
		&BadTerm{"entirely eliminate", "'%s' is redundant."},
		&BadTerm{"most especially", "'%s' is redundant."},
		&BadTerm{"actual fact", "'%s' is redundant."},
		&BadTerm{"true facts", "'%s' is redundant."},
		&BadTerm{"future forecast", "'%s' is redundant."},
		&BadTerm{"founding forefathers", "'%s' is redundant."},
		&BadTerm{"free and gratis", "'%s' is redundant."},
		&BadTerm{"free gratis", "'%s' is redundant."},
		&BadTerm{"completely full", "'%s' is redundant."},
		&BadTerm{"basic fundamentals", "'%s' is redundant."},
		&BadTerm{"free gift", "'%s' is redundant."},
		&BadTerm{"new innovation", "'%s' is redundant."},
		&BadTerm{"interact with each other", "'%s' is redundant."},
		&BadTerm{"large-size", "'%s' is redundant."},
		&BadTerm{"meld together", "'%s' is redundant."},
		&BadTerm{"merge together", "'%s' is redundant."},
		&BadTerm{"mingle together", "'%s' is redundant."},
		&BadTerm{"mix together", "'%s' is redundant."},
		&BadTerm{"mutual feelings for each other", "'%s' is redundant."},
		&BadTerm{"mutual respect for each other", "'%s' is redundant."},
		&BadTerm{"absolute necessity", "'%s' is redundant."},
		&BadTerm{"blatantly obvious", "'%s' is redundant."},
		&BadTerm{"pause for a moment", "'%s' is redundant."},
		&BadTerm{"advance planning", "'%s' is redundant."},
		&BadTerm{"future plans", "'%s' is redundant."},
		&BadTerm{"pooled together", "'%s' is redundant."},
		&BadTerm{"potable drinking water", "'%s' is redundant."},
		&BadTerm{"potable drinking water", "'%s' is redundant."},
		&BadTerm{"new recruit", "'%s' is redundant."},
		&BadTerm{"reelected for another term", "'%s' is redundant."},
		&BadTerm{"refer back", "'%s' is redundant."},
		&BadTerm{"regress back", "'%s' is redundant."},
		&BadTerm{"repay them back", "'%s' is redundant."},
		&BadTerm{"repay back", "'%s' is redundant."},
		&BadTerm{"repeat again", "'%s' is redundant."},
		&BadTerm{"repeat back", "'%s' is redundant."},
		&BadTerm{"repeat the same", "'%s' is redundant."},
		&BadTerm{"repeated the same", "'%s' is redundant."},
		&BadTerm{"temporary reprieve", "'%s' is redundant."},
		&BadTerm{"brief respite", "'%s' is redundant."},
		&BadTerm{"retreat back", "'%s' is redundant."},
		&BadTerm{"return back", "'%s' is redundant."},
		&BadTerm{"closely scrutinize", "'%s' is redundant."},
		&BadTerm{"software program", "'%s' is redundant."},
		&BadTerm{"surrounded on all sides", "'%s' is redundant."},
		&BadTerm{"the whole entire nation", "'%s' is redundant."},
		&BadTerm{"throughout the entire", "'%s' is redundant."},
		&BadTerm{"timpani drum", "'%s' is redundant."},
		&BadTerm{"pair of twins", "'%s' is redundant."},
		&BadTerm{"unfilled vacancy", "'%s' is redundant."},
		&BadTerm{"various different", "'%s' is redundant."},
		&BadTerm{"former veteran", "'%s' is redundant."},
		&BadTerm{"visible to the eye", "'%s' is redundant."},
		&BadTerm{"professional vocation", "'%s' is redundant."},
		&BadTerm{"while at the same time", "'%s' is redundant."},
		&BadTerm{"absolutely essential", "'%s' is redundant."},
		&BadTerm{"absolutely necessary", "'%s' is redundant."},
		&BadTerm{"bo staff", "'%s' is redundant."},
		&BadTerm{"challah bread", "'%s' is redundant."},
		&BadTerm{"hallah bread", "'%s' is redundant."},
		&BadTerm{"challah bread", "'%s' is redundant."},
		&BadTerm{"i myself", "'%s' is redundant."},
		&BadTerm{"i personally", "'%s' is redundant."},
		&BadTerm{"mount fujiyama", "'%s' is redundant."},
		&BadTerm{"milky way galaxy", "'%s' is redundant."},
		&BadTerm{"rio grande river", "'%s' is redundant."},
		&BadTerm{"old adage", "'%s' is redundant."},
		&BadTerm{"add a further", "'%s' is redundant."},
		&BadTerm{"add an additional", "'%s' is redundant."},
		&BadTerm{"advance forward", "'%s' is redundant."},
		&BadTerm{"alternative choice", "'%s' is redundant."},
		&BadTerm{"amaretto almond", "'%s' is redundant."},
		&BadTerm{"completely annihilate", "'%s' is redundant."},
		&BadTerm{"annual anniversary", "'%s' is redundant."},
		&BadTerm{"unnamed anonymous", "'%s' is redundant."},
		&BadTerm{"equally as", "'%s' is redundant."},
		&BadTerm{"ascend up", "'%s' is redundant."},
		&BadTerm{"ask the question", "'%s' is redundant."},
		&BadTerm{"assemble together", "'%s' is redundant."},
		&BadTerm{"at the present time the", "'%s' is redundant."},
		&BadTerm{"at this point in time", "'%s' is redundant."},
		&BadTerm{"attach together", "'%s' is redundant."},
		&BadTerm{"autumn season", "'%s' is redundant."},
		&BadTerm{"bald-headed", "'%s' is redundant."},
		&BadTerm{"balsa wood", "'%s' is redundant."},
		&BadTerm{"personal belongings", "'%s' is redundant."},
		&BadTerm{"desirable benefits", "'%s' is redundant."},
		&BadTerm{"bento box", "'%s' is redundant."},
		&BadTerm{"best ever", "'%s' is redundant."},
		&BadTerm{"tiny bit", "'%s' is redundant."},
		&BadTerm{"blend together", "'%s' is redundant."},
		&BadTerm{"common bond", "'%s' is redundant."},
		&BadTerm{"added bonus", "'%s' is redundant."},
		&BadTerm{"extra bonus", "'%s' is redundant."},
		&BadTerm{"bouquet of flowers", "'%s' is redundant."},
		&BadTerm{"major breakthrough", "'%s' is redundant."},
		&BadTerm{"new bride", "'%s' is redundant."},
		&BadTerm{"brief in duration", "'%s' is redundant."},
		&BadTerm{"bruin bear", "'%s' is redundant."},
		&BadTerm{"burning hot", "'%s' is redundant."},
		&BadTerm{"cacophony of sound", "'%s' is redundant."},
		&BadTerm{"brief cameo", "'%s' is redundant."},
		&BadTerm{"cameo appearance", "'%s' is redundant."},
		&BadTerm{"cancel out", "'%s' is redundant."},
		&BadTerm{"cash money", "'%s' is redundant."},
		&BadTerm{"chai tea", "'%s' is redundant."},
		&BadTerm{"random chance", "'%s' is redundant."},
		&BadTerm{"personal charm", "'%s' is redundant."},
		&BadTerm{"circle around", "'%s' is redundant."},
		&BadTerm{"round circle", "'%s' is redundant."},
		&BadTerm{"circulate around", "'%s' is redundant."},
		&BadTerm{"classify into groups", "'%s' is redundant."},
		&BadTerm{"fellow classmates", "'%s' is redundant."},
		&BadTerm{"old cliche", "'%s' is redundant."},
		&BadTerm{"old cliche", "'%s' is redundant."},
		&BadTerm{"climb up", "'%s' is redundant."},
		&BadTerm{"time clock", "'%s' is redundant."},
		&BadTerm{"collaborate together", "'%s' is redundant."},
		&BadTerm{"joint collaboration", "'%s' is redundant."},
		&BadTerm{"fellow colleague", "'%s' is redundant."},
		&BadTerm{"combine together", "'%s' is redundant."},
		&BadTerm{"commute back and forth", "'%s' is redundant."},
		&BadTerm{"compete with each other", "'%s' is redundant."},
		&BadTerm{"comprise of", "'%s' is redundant."},
		&BadTerm{"comprises of", "'%s' is redundant."},
		&BadTerm{"first conceived", "'%s' is redundant."},
		&BadTerm{"final conclusion", "'%s' is redundant."},
		&BadTerm{"confer together", "'%s' is redundant."},
		&BadTerm{"direct confrontation", "'%s' is redundant."},
		&BadTerm{"connect together", "'%s' is redundant."},
		&BadTerm{"connect up", "'%s' is redundant."},
		&BadTerm{"consensus of opinion", "'%s' is redundant."},
		&BadTerm{"general consensus", "'%s' is redundant."},
		&BadTerm{"consult with", "'%s' is redundant."},
		&BadTerm{"oral conversation", "'%s' is redundant."},
		&BadTerm{"cool down", "'%s' is redundant."},
		&BadTerm{"cooperate together", "'%s' is redundant."},
		&BadTerm{"mutual cooperation", "'%s' is redundant."},
		&BadTerm{"duplicate copy", "'%s' is redundant."},
		&BadTerm{"inner core", "'%s' is redundant."},
		&BadTerm{"cost the sum of", "'%s' is redundant."},
		&BadTerm{"could possibly", "'%s' is redundant."},
		&BadTerm{"money-saving coupon", "'%s' is redundant."},
		&BadTerm{"originally created", "'%s' is redundant."},
		&BadTerm{"crisis situation", "'%s' is redundant."},
		&BadTerm{"crouch down", "'%s' is redundant."},
		&BadTerm{"now currently", "'%s' is redundant."},
		&BadTerm{"old custom", "'%s' is redundant."},
		&BadTerm{"usual custom", "'%s' is redundant."},
		&BadTerm{"serious danger", "'%s' is redundant."},
		&BadTerm{"dates back", "'%s' is redundant."},
		&BadTerm{"definite decision", "'%s' is redundant."},
		&BadTerm{"depreciate in value", "'%s' is redundant."},
		&BadTerm{"descend down", "'%s' is redundant."},
		&BadTerm{"totally destroy", "'%s' is redundant."},
		&BadTerm{"completely destroyed", "'%s' is redundant."},
		&BadTerm{"total destruction", "'%s' is redundant."},
		&BadTerm{"specific details", "'%s' is redundant."},
		&BadTerm{"difficult dilemma", "'%s' is redundant."},
		&BadTerm{"disappear from sight", "'%s' is redundant."},
		&BadTerm{"originally discovered", "'%s' is redundant."},
		&BadTerm{"dive down", "'%s' is redundant."},
		&BadTerm{"over and done with", "'%s' is redundant."},
		&BadTerm{"illustrated drawing", "'%s' is redundant."},
		&BadTerm{"drop down", "'%s' is redundant."},
		&BadTerm{"sand dune", "'%s' is redundant."},
		&BadTerm{"during the course of", "'%s' is redundant."},
		&BadTerm{"dwindle down", "'%s' is redundant."},
		&BadTerm{"dwindled down", "'%s' is redundant."},
		&BadTerm{"each and every", "'%s' is redundant."},
		&BadTerm{"earlier in time", "'%s' is redundant."},
		&BadTerm{"completely eliminate", "'%s' is redundant."},
		&BadTerm{"eliminate altogether", "'%s' is redundant."},
		&BadTerm{"entirely eliminate", "'%s' is redundant."},
		&BadTerm{"glowing ember", "'%s' is redundant."},
		&BadTerm{"burning embers", "'%s' is redundant."},
		&BadTerm{"emergency situation", "'%s' is redundant."},
		&BadTerm{"unexpected emergency", "'%s' is redundant."},
		&BadTerm{"empty out", "'%s' is redundant."},
		&BadTerm{"enclosed herein", "'%s' is redundant."},
		&BadTerm{"final end", "'%s' is redundant."},
		&BadTerm{"completely engulfed", "'%s' is redundant."},
		&BadTerm{"enter in", "'%s' is redundant."},
		&BadTerm{"enter into", "'%s' is redundant."},
		&BadTerm{"equal to one another", "'%s' is redundant."},
		&BadTerm{"eradicate completely", "'%s' is redundant."},
		&BadTerm{"absolutely essential", "'%s' is redundant."},
		&BadTerm{"estimated at about", "'%s' is redundant."},
		&BadTerm{"estimated at approximately", "'%s' is redundant."},
		&BadTerm{"estimated at around", "'%s' is redundant."},
		&BadTerm{"and etc.", "'%s' is redundant."},
		&BadTerm{"evolve over time", "'%s' is redundant."},
		&BadTerm{"over exaggerate", "'%s' is redundant."},
		&BadTerm{"exited from", "'%s' is redundant."},
		&BadTerm{"actual experience", "'%s' is redundant."},
		&BadTerm{"past experience", "'%s' is redundant."},
		&BadTerm{"knowledgeable experts", "'%s' is redundant."},
		&BadTerm{"extradite back", "'%s' is redundant."},
		&BadTerm{"face up to the consequences", "'%s' is redundant."},
		&BadTerm{"face up to the fact", "'%s' is redundant."},
		&BadTerm{"face up to the challenge", "'%s' is redundant."},
		&BadTerm{"face up to the problem", "'%s' is redundant."},
		&BadTerm{"facilitate easier", "'%s' is redundant."},
		&BadTerm{"established fact", "'%s' is redundant."},
		&BadTerm{"actual facts", "'%s' is redundant."},
		&BadTerm{"hard facts", "'%s' is redundant."},
		&BadTerm{"true facts", "'%s' is redundant."},
		&BadTerm{"passing fad", "'%s' is redundant."},
		&BadTerm{"fall down", "'%s' is redundant."},
		&BadTerm{"fall season", "'%s' is redundant."},
		&BadTerm{"major feat", "'%s' is redundant."},
		&BadTerm{"feel inside", "'%s' is redundant."},
		&BadTerm{"inner feelings", "'%s' is redundant."},
		&BadTerm{"few in number", "'%s' is redundant."},
		&BadTerm{"completely filled", "'%s' is redundant."},
		&BadTerm{"filled to capacity", "'%s' is redundant."},
		&BadTerm{"first of all", "'%s' is redundant."},
		&BadTerm{"first time ever", "'%s' is redundant."},
		&BadTerm{"closed fist", "'%s' is redundant."},
		&BadTerm{"fly through the air", "'%s' is redundant."},
		&BadTerm{"main focus", "'%s' is redundant."},
		&BadTerm{"focus in", "'%s' is redundant."},
		&BadTerm{"follow after", "'%s' is redundant."},
		&BadTerm{"as for example", "'%s' is redundant."},
		&BadTerm{"forever and ever", "'%s' is redundant."},
		&BadTerm{"for free", "'%s' is redundant."},
		&BadTerm{"personal friend", "'%s' is redundant."},
		&BadTerm{"personal friendship", "'%s' is redundant."},
		&BadTerm{"full to capacity", "'%s' is redundant."},
		&BadTerm{"basic fundamentals", "'%s' is redundant."},
		&BadTerm{"fuse together", "'%s' is redundant."},
		&BadTerm{"gather up", "'%s' is redundant."},
		&BadTerm{"gather together", "'%s' is redundant."},
		&BadTerm{"get up on (his|your) feet", "'%s' is redundant."},
		&BadTerm{"free gift", "'%s' is redundant."},
		&BadTerm{"free gifts", "'%s' is redundant."},
		&BadTerm{"ultimate goal", "'%s' is redundant."},
		&BadTerm{"grow in size", "'%s' is redundant."},
		&BadTerm{"absolute guarantee", "'%s' is redundant."},
		&BadTerm{"armed gunman", "'%s' is redundant."},
		&BadTerm{"armed gunmen", "'%s' is redundant."},
		&BadTerm{"native habitat", "'%s' is redundant."},
		&BadTerm{"had done previously", "'%s' is redundant."},
		&BadTerm{"two equal halves", "'%s' is redundant."},
		&BadTerm{"safe haven", "'%s' is redundant."},
		&BadTerm{"heat up", "'%s' is redundant."},
		&BadTerm{"past history", "'%s' is redundant."},
		&BadTerm{"hoist up", "'%s' is redundant."},
		&BadTerm{"empty hole", "'%s' is redundant."},
		&BadTerm{"head honcho", "'%s' is redundant."},
		&BadTerm{"frozen ice", "'%s' is redundant."},
		&BadTerm{"perfect ideal", "'%s' is redundant."},
		&BadTerm{"same identical", "'%s' is redundant."},
		&BadTerm{"positive identification", "'%s' is redundant."},
		&BadTerm{"foreign imports", "'%s' is redundant."},
		&BadTerm{"sudden impulse", "'%s' is redundant."},
		&BadTerm{"in actual fact", "'%s' is redundant."},
		&BadTerm{"outside in the yard", "'%s' is redundant."},
		&BadTerm{"all inclusive", "'%s' is redundant."},
		&BadTerm{"incredible to believe", "'%s' is redundant."},
		&BadTerm{"present incumbent", "'%s' is redundant."},
		&BadTerm{"private industry", "'%s' is redundant."},
		&BadTerm{"harmful injuries", "'%s' is redundant."},
		&BadTerm{"new innovation", "'%s' is redundant."},
		&BadTerm{"innovative new", "'%s' is redundant."},
		&BadTerm{"new innovative", "'%s' is redundant."},
		&BadTerm{"natural instinct", "'%s' is redundant."},
		&BadTerm{"integrate together", "'%s' is redundant."},
		&BadTerm{"integrate with each other", "'%s' is redundant."},
		&BadTerm{"interdependent on each other", "'%s' is redundant."},
		&BadTerm{"mutually interdependent", "'%s' is redundant."},
		&BadTerm{"introduced for the first time", "'%s' is redundant."},
		&BadTerm{"new invention", "'%s' is redundant."},
		&BadTerm{"kneel down", "'%s' is redundant."},
		&BadTerm{"knots per hour", "'%s' is redundant."},
		&BadTerm{"lift up", "'%s' is redundant."},
		&BadTerm{"still lingers", "'%s' is redundant."},
		&BadTerm{"look ahead to the future", "'%s' is redundant."},
		&BadTerm{"three-way love triangle", "'%s' is redundant."},
		&BadTerm{"constantly maintained", "'%s' is redundant."},
		&BadTerm{"manually by hand", "'%s' is redundant."},
		&BadTerm{"boat marina", "'%s' is redundant."},
		&BadTerm{"may possibly", "'%s' is redundant."},
		&BadTerm{"meet with each other", "'%s' is redundant."},
		&BadTerm{"meet together", "'%s' is redundant."},
		&BadTerm{"past memories", "'%s' is redundant."},
		&BadTerm{"merge together", "'%s' is redundant."},
		&BadTerm{"merged together", "'%s' is redundant."},
		&BadTerm{"meshed together", "'%s' is redundant."},
		&BadTerm{"twelve midnight", "'%s' is redundant."},
		&BadTerm{"migraine headache", "'%s' is redundant."},
		&BadTerm{"minestrone soup", "'%s' is redundant."},
		&BadTerm{"mix together", "'%s' is redundant."},
		&BadTerm{"brief moment", "'%s' is redundant."},
		&BadTerm{"moment in time", "'%s' is redundant."},
		&BadTerm{"complete monopoly", "'%s' is redundant."},
		&BadTerm{"wall mural", "'%s' is redundant."},
		&BadTerm{"mutual respect for each other", "'%s' is redundant."},
		&BadTerm{"mutually dependent on each other", "'%s' is redundant."},
		&BadTerm{"unsolved mystery", "'%s' is redundant."},
		&BadTerm{"nape of her neck", "'%s' is redundant."},
		&BadTerm{"absolutely necessary", "'%s' is redundant."},
		&BadTerm{"never at any time", "'%s' is redundant."},
		&BadTerm{"12 noon", "'%s' is redundant."},
		&BadTerm{"12 o'clock noon", "'%s' is redundant."},
		&BadTerm{"high noon", "'%s' is redundant."},
		&BadTerm{"twelve noon", "'%s' is redundant."},
		&BadTerm{"nostalgia for the past", "'%s' is redundant."},
		&BadTerm{"number of different", "'%s' is redundant."},
		&BadTerm{"exposed opening", "'%s' is redundant."},
		&BadTerm{"my personal opinion", "'%s' is redundant."},
		&BadTerm{"exact opposites?", "'%s' is redundant."},
		&BadTerm{"polar opposites?", "'%s' is redundant."},
		&BadTerm{"orbits around", "'%s' is redundant."},
		&BadTerm{"final outcome", "'%s' is redundant."},
		&BadTerm{"universal panacea", "'%s' is redundant."},
		&BadTerm{"now pending", "'%s' is redundant."},
		&BadTerm{"penetrate through", "'%s' is redundant."},
		&BadTerm{"still persists", "'%s' is redundant."},
		&BadTerm{"old pioneer", "'%s' is redundant."},
		&BadTerm{"proposed plan", "'%s' is redundant."},
		&BadTerm{"plan ahead", "'%s' is redundant."},
		&BadTerm{"plan in advance", "'%s' is redundant."},
		&BadTerm{"advance planning", "'%s' is redundant."},
		&BadTerm{"forward planning", "'%s' is redundant."},
		&BadTerm{"future plans?", "'%s' is redundant."},
		&BadTerm{"point in time", "'%s' is redundant."},
		&BadTerm{"sharp point", "'%s' is redundant."},
		&BadTerm{"postpone until later", "'%s' is redundant."},
		&BadTerm{"pouring down rain", "'%s' is redundant."},
		&BadTerm{"advance preview", "'%s' is redundant."},
		&BadTerm{"previously listed above", "'%s' is redundant."},
		&BadTerm{"probed into", "'%s' is redundant."},
		&BadTerm{"proceed ahead", "'%s' is redundant."},
		&BadTerm{"artificial prosthesis", "'%s' is redundant."},
		&BadTerm{"old proverb", "'%s' is redundant."},
		&BadTerm{"put off until later", "'%s' is redundant."},
		&BadTerm{"re-elect for another term", "'%s' is redundant."},
		&BadTerm{"reason is because", "'%s' is redundant."},
		&BadTerm{"recur again", "'%s' is redundant."},
		&BadTerm{"future recurrence", "'%s' is redundant."},
		&BadTerm{"refer back", "'%s' is redundant."},
		&BadTerm{"reflect back", "'%s' is redundant."},
		&BadTerm{"continue to remain", "'%s' is redundant."},
		&BadTerm{"still remains", "'%s' is redundant."},
		&BadTerm{"exact replica", "'%s' is redundant."},
		&BadTerm{"reply back", "'%s' is redundant."},
		&BadTerm{"advance reservations", "'%s' is redundant."},
		&BadTerm{"retreat back", "'%s' is redundant."},
		&BadTerm{"revert back", "'%s' is redundant."},
		&BadTerm{"round in shape", "'%s' is redundant."},
		&BadTerm{"rough rule of thumb", "'%s' is redundant."},
		&BadTerm{"unconfirmed rumor", "'%s' is redundant."},
		&BadTerm{"rustic country", "'%s' is redundant."},
		&BadTerm{"same exact", "'%s' is redundant."},
		&BadTerm{"exact same", "'%s' is redundant."},
		&BadTerm{"precise same", "'%s' is redundant."},
		&BadTerm{"precisely the same", "'%s' is redundant."},
		&BadTerm{"safe sanctuary", "'%s' is redundant."},
		&BadTerm{"full satisfaction", "'%s' is redundant."},
		&BadTerm{"scrutinize in detail", "'%s' is redundant."},
		&BadTerm{"careful scrutiny", "'%s' is redundant."},
		&BadTerm{"close scrutiny", "'%s' is redundant."},
		&BadTerm{"secret that cannot be told", "'%s' is redundant."},
		&BadTerm{"seek to find", "'%s' is redundant."},
		&BadTerm{"separated apart from each other", "'%s' is redundant."},
		&BadTerm{"share together", "'%s' is redundant."},
		&BadTerm{"shiny in appearance", "'%s' is redundant."},
		&BadTerm{"truly sincere", "'%s' is redundant."},
		&BadTerm{"sink down", "'%s' is redundant."},
		&BadTerm{"skipped over", "'%s' is redundant."},
		&BadTerm{"soft in texture", "'%s' is redundant."},
		&BadTerm{"soft to the touch", "'%s' is redundant."},
		&BadTerm{"sole of the foot", "'%s' is redundant."},
		&BadTerm{"some time to come", "'%s' is redundant."},
		&BadTerm{"small speck", "'%s' is redundant."},
		&BadTerm{"rate of speed", "'%s' is redundant."},
		&BadTerm{"spell out in detail", "'%s' is redundant."},
		&BadTerm{"spiked upwards?", "'%s' is redundant."},
		&BadTerm{"spring season", "'%s' is redundant."},
		&BadTerm{"anonymous stranger", "'%s' is redundant."},
		&BadTerm{"live studio audience", "'%s' is redundant."},
		&BadTerm{"underground subway", "'%s' is redundant."},
		&BadTerm{"sufficient enough", "'%s' is redundant."},
		&BadTerm{"summer season", "'%s' is redundant."},
		&BadTerm{"absolutely sure", "'%s' is redundant."},
		&BadTerm{"unexpected surprise", "'%s' is redundant."},
		&BadTerm{"completely surround", "'%s' is redundant."},
		&BadTerm{"surrounded on all sides", "'%s' is redundant."},
		&BadTerm{"tall in height", "'%s' is redundant."},
		&BadTerm{"tall in stature", "'%s' is redundant."},
		&BadTerm{"mental telepathy", "'%s' is redundant."},
		&BadTerm{"ten in number", "'%s' is redundant."},
		&BadTerm{"these ones", "'%s' is redundant."},
		&BadTerm{"those ones", "'%s' is redundant."},
		&BadTerm{"open trench", "'%s' is redundant."},
		&BadTerm{"honest truth", "'%s' is redundant."},
		&BadTerm{"frozen tundra", "'%s' is redundant."},
		&BadTerm{"final ultimatum", "'%s' is redundant."},
		&BadTerm{"undergraduate student", "'%s' is redundant."},
		&BadTerm{"vacillate back and forth", "'%s' is redundant."},
		&BadTerm{"former veteran", "'%s' is redundant."},
		&BadTerm{"visible to the eye", "'%s' is redundant."},
		&BadTerm{"warn in advance", "'%s' is redundant."},
		&BadTerm{"advance warning", "'%s' is redundant."},
		&BadTerm{"hot water heater", "'%s' is redundant."},
		&BadTerm{"in which we live in", "'%s' is redundant."},
		&BadTerm{"winter season", "'%s' is redundant."},
		&BadTerm{"live witness", "'%s' is redundant."},
		&BadTerm{"yakitori chicken", "'%s' is redundant."},
		&BadTerm{"yerba mate tea", "'%s' is redundant."},
		&BadTerm{"affirmative yes", "'%s' is redundant."},
	}
}
