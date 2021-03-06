package main

// NEEDCAPS

// ShouldNotBias returns a slice of BadTerm's, none of which should be in the
// a text (case insensitive).
func ShouldNotBias() []TextCheck {
	return []TextCheck{
		// Ablist
		&BadTerm{"downs syndrome", "'%s' is incorrect. You mean Down Syndrome."},
		&BadTerm{"ablebodied", "'%s' is ablist."},
		&BadTerm{"addicts?", "'%s' is ablist."},
		&BadTerm{"afflicted with a disability", "'%s' is ablist."},
		&BadTerm{"afflicted with a intellectual disability", "'%s' is ablist."},
		&BadTerm{"afflicted with a polio", "'%s' is ablist."},
		&BadTerm{"afflicted with aids", "'%s' is ablist."},
		&BadTerm{"afflicted with an injury", "'%s' is ablist."},
		&BadTerm{"afflicted with disabilities", "'%s' is ablist."},
		&BadTerm{"afflicted with injuries", "'%s' is ablist."},
		&BadTerm{"afflicted with intellectual disabilities", "'%s' is ablist."},
		&BadTerm{"afflicted with multiple sclerosis", "'%s' is ablist."},
		&BadTerm{"afflicted with polio", "'%s' is ablist."},
		&BadTerm{"afflicted with psychosis", "'%s' is ablist."},
		&BadTerm{"afflicted with schizophrenia", "'%s' is ablist."},
		&BadTerm{"aids victim", "'%s' is ablist."},
		&BadTerm{"alcoholic", "'%s' is ablist."},
		&BadTerm{"amputee", "'%s' is ablist."},
		&BadTerm{"asylum", "'%s' is ablist."},
		&BadTerm{"autistic", "'%s' is ablist."},
		&BadTerm{"barren", "'%s' is ablist."},
		&BadTerm{"batshit", "'%s' is ablist."},
		&BadTerm{"bedlam", "'%s' is ablist."},
		&BadTerm{"bipolar", "'%s' is ablist."},
		&BadTerm{"birth defect", "'%s' is ablist."},
		&BadTerm{"blind eye to", "'%s' is ablist."},
		&BadTerm{"blind to", "'%s' is ablist."},
		&BadTerm{"blinded by", "'%s' is ablist."},
		&BadTerm{"bound to a wheelchair", "'%s' is ablist."},
		&BadTerm{"cleftlipped", "'%s' is ablist."},
		&BadTerm{"confined to a wheelchair", "'%s' is ablist."},
		&BadTerm{"crazy", "'%s' is ablist."},
		&BadTerm{"cretin", "'%s' is ablist."},
		&BadTerm{"cripple", "'%s' is ablist."},
		&BadTerm{"crippled", "'%s' is ablist."},
		&BadTerm{"daft", "'%s' is ablist."},
		&BadTerm{"deaf and dumb", "'%s' is ablist."},
		&BadTerm{"deaf ear to", "'%s' is ablist."},
		&BadTerm{"deaf to", "'%s' is ablist."},
		&BadTerm{"deafened by", "'%s' is ablist."},
		&BadTerm{"deafmute", "'%s' is ablist."},
		&BadTerm{"demented", "'%s' is ablist."},
		&BadTerm{"depressed", "'%s' is ablist."},
		&BadTerm{"detox", "'%s' is ablist."},
		&BadTerm{"detox center", "'%s' is ablist."},
		&BadTerm{"disabled", "'%s' is ablist."},
		&BadTerm{"dumb", "'%s' is ablist."},
		&BadTerm{"dyslexic", "'%s' is ablist."},
		&BadTerm{"epileptic", "'%s' is ablist."},
		&BadTerm{"family burden", "'%s' is ablist."},
		&BadTerm{"feeble minded", "'%s' is ablist."},
		&BadTerm{"feebleminded", "'%s' is ablist."},
		&BadTerm{"handicapped", "'%s' is ablist."},
		&BadTerm{"handicapped parking", "'%s' is ablist."},
		&BadTerm{"harelip", "'%s' is ablist."},
		&BadTerm{"harelipped", "'%s' is ablist."},
		&BadTerm{"hearing impaired", "'%s' is ablist."},
		&BadTerm{"hearing impairment", "'%s' is ablist."},
		&BadTerm{"idiot", "'%s' is ablist."},
		&BadTerm{"imbecile", "'%s' is ablist."},
		&BadTerm{"infantile paralysis", "'%s' is ablist."},
		&BadTerm{"insane", "'%s' is ablist."},
		&BadTerm{"insanity", "'%s' is ablist."},
		&BadTerm{"insomniacs?", "'%s' is ablist."},
		&BadTerm{"intellectually disabled", "'%s' is ablist."},
		&BadTerm{"intellectually disabled people", "'%s' is ablist."},
		&BadTerm{"lame", "'%s' is ablist."},
		&BadTerm{"learning disabled", "'%s' is ablist."},
		&BadTerm{"loony", "'%s' is ablist."},
		&BadTerm{"loony bin", "'%s' is ablist."},
		&BadTerm{"lunacy", "'%s' is ablist."},
		&BadTerm{"lunatic", "'%s' is ablist."},
		&BadTerm{"madhouse", "'%s' is ablist."},
		&BadTerm{"maniac", "'%s' is ablist."},
		&BadTerm{"mental defective", "'%s' is ablist."},
		&BadTerm{"mentally ill", "'%s' is ablist."},
		&BadTerm{"midget", "'%s' is ablist."},
		&BadTerm{"mongoloid", "'%s' is ablist."},
		&BadTerm{"moron", "'%s' is ablist."},
		&BadTerm{"moronic", "'%s' is ablist."},
		&BadTerm{"multiple sclerosis victim", "'%s' is ablist."},
		&BadTerm{"nuts", "'%s' is ablist."},
		&BadTerm{`o\.c\.d`, "'%s' is ablist."},
		&BadTerm{`o\.c\.d\.`, "'%s' is ablist."},
		&BadTerm{"ocd", "'%s' is ablist."},
		&BadTerm{"paraplegic", "'%s' is ablist."},
		&BadTerm{"psycho", "'%s' is ablist."},
		&BadTerm{"psychopathology", "'%s' is ablist."},
		&BadTerm{"psychotic", "'%s' is ablist."},
		&BadTerm{"quadriplegic", "'%s' is ablist."},
		&BadTerm{"rehab", "'%s' is ablist."},
		&BadTerm{"rehab center", "'%s' is ablist."},
		&BadTerm{"restricted to a wheelchair", "'%s' is ablist."},
		&BadTerm{"retarded", "'%s' is ablist."},
		&BadTerm{"retards?", "'%s' is ablist."},
		&BadTerm{"schizo", "'%s' is ablist."},
		&BadTerm{"schizophrenic", "'%s' is ablist."},
		&BadTerm{"senile", "'%s' is ablist."},
		&BadTerm{"simpleton", "'%s' is ablist."},
		&BadTerm{"sociopath", "'%s' is ablist."},
		&BadTerm{"sociopaths", "'%s' is ablist."},
		&BadTerm{"spastic", "'%s' is ablist."},
		&BadTerm{"spaz", "'%s' is ablist."},
		&BadTerm{"special olympians", "'%s' is ablist."},
		&BadTerm{"special olympic athletes", "'%s' is ablist."},
		&BadTerm{"stammering", "'%s' is ablist."},
		&BadTerm{"stroke victim", "'%s' is ablist."},
		&BadTerm{"stupid", "'%s' is ablist."},
		&BadTerm{"stutterer", "'%s' is ablist."},
		&BadTerm{"suffer from aids", "'%s' is ablist."},
		&BadTerm{"suffer from an injury", "'%s' is ablist."},
		&BadTerm{"suffer from injuries", "'%s' is ablist."},
		&BadTerm{"suffering from a disability", "'%s' is ablist."},
		&BadTerm{"suffering from a polio", "'%s' is ablist."},
		&BadTerm{"suffering from a stroke", "'%s' is ablist."},
		&BadTerm{"suffering from aids", "'%s' is ablist."},
		&BadTerm{"suffering from an injury", "'%s' is ablist."},
		&BadTerm{"suffering from an intellectual disability", "'%s' is ablist."},
		&BadTerm{"suffering from disabilities", "'%s' is ablist."},
		&BadTerm{"suffering from injuries", "'%s' is ablist."},
		&BadTerm{"suffering from intellectual disabilities", "'%s' is ablist."},
		&BadTerm{"suffering from multiple sclerosis", "'%s' is ablist."},
		&BadTerm{"suffering from polio", "'%s' is ablist."},
		&BadTerm{"suffering from psychosis", "'%s' is ablist."},
		&BadTerm{"suffering from schizophrenia", "'%s' is ablist."},
		&BadTerm{"suffers from aids", "'%s' is ablist."},
		&BadTerm{"suffers from an injury", "'%s' is ablist."},
		&BadTerm{"suffers from disabilities", "'%s' is ablist."},
		&BadTerm{"suffers from injuries", "'%s' is ablist."},
		&BadTerm{"suffers from intellectual disabilities", "'%s' is ablist."},
		&BadTerm{"suffers from multiple sclerosis", "'%s' is ablist."},
		&BadTerm{"suffers from polio", "'%s' is ablist."},
		&BadTerm{"suffers from psychosis", "'%s' is ablist."},
		&BadTerm{"suffers from schizophrenia", "'%s' is ablist."},
		&BadTerm{"tourettes disorder", "'%s' is ablist."},
		&BadTerm{"tourettes syndrome", "'%s' is ablist."},
		&BadTerm{"vertically challenged", "'%s' is ablist."},
		&BadTerm{"victim of a stroke", "'%s' is ablist."},
		&BadTerm{"victim of aids", "'%s' is ablist."},
		&BadTerm{"victim of an injury", "'%s' is ablist."},
		&BadTerm{"victim of injuries", "'%s' is ablist."},
		&BadTerm{"victim of multiple sclerosis", "'%s' is ablist."},
		&BadTerm{"victim of polio", "'%s' is ablist."},
		&BadTerm{"victim of psychosis", "'%s' is ablist."},
		&BadTerm{"wacko", "'%s' is ablist."},
		&BadTerm{"whacko", "'%s' is ablist."},
		&BadTerm{"wheelchair bound", "'%s' is ablist."},

		// LGBTQ
		&BadTerm{"bathroom bill", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"biologically female", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"biologically male", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"born a man", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"born a woman", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"dyke", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"fag", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"faggot", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"gay agenda", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"gay lifestyle", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"gay rights", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"genetically female", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"genetically male", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"hermaphrodite", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"hermaphroditic", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"heshe", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"homo", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"homosexual", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"homosexual agenda", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"homosexual couple", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"homosexual lifestyle", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"homosexual marriage", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"homosexual relations", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"homosexual relationship", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"pseudo hermaphrodite", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"pseudo hermaphroditic", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"pseudohermaphrodite", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"pseudohermaphroditic", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"sex change", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"sexchange", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"sexual preference", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"she male", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"shehe", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"shemale", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"sodomite", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"tranny", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"transgendered", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"transgenderism", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"transgenders", "'%s' is inconsiderate. Check yourself."},
		&BadTerm{"transvestite", "'%s' is inconsiderate. Check yourself."},

		// Muslim
		&BadTerm{"islamists?", "'%s' is inconsiderate. Use either Muslims (or people of Islamic faith) OR fanatic/zealot depending on your meaning."},

		// Race
		&BadTerm{"masters?", "'%s' is inappropriate. Consider primary/replica, hub/spoke, or some other metaphor."},
		&BadTerm{"slaves?", "'%s' is inappropriate. Consider primary/replica, hub/spoke, or some other metaphor."},
		&BadTerm{"eskimos?", "'%s' is inappropriate. You mean Inuit"},
		&BadTerm{"orientals?", "'%s' is inappropriate. You mean Asian"},
		&BadTerm{"nonwhite", "'%s' is inappropriate. You mean person/people of color."},
		&BadTerm{"non white", "'%s' is inappropriate. You mean person/people of color."},
		&BadTerm{"ghetto", "'%s' is inappropriate. Use a noun like projects or an adjective like urban."},

		// Suicide
		&BadTerm{"commit suicide", "'%s' is inconsiderate. Consider 'die by suicide'."},
		&BadTerm{"committed suicide", "'%s' is inconsiderate. Consider 'died by suicide'."},
		&BadTerm{"suicide epidemic", "'%s' is inconsiderate. Consider 'rise in suicides'."},
	}
}
