package jpush

type Audience struct {
	IsAll bool
	Value map[string][]string
}

func NewAudience() *Audience {
	return &Audience{
		Value: make(map[string][]string),
	}
}

func (a *Audience) Interface() interface{} {
	if a.IsAll {
		return "all"
	}
	return a.Value
}

func (a *Audience) All() {
	a.IsAll = true
}

func (a *Audience) SetTag(tags ...string) {
	a.set("tag", tags)
}

func (a *Audience) SetTagAnd(tagAnds ...string) {
	a.set("tag_and", tagAnds)
}

func (a *Audience) SetTagNot(tagNots ...string) {
	a.set("tag_not", tagNots)
}

func (a *Audience) SetRegistrationId(regIds ...string) {
	a.set("registration_id", regIds)
}

func (a *Audience) SetSegment(segments ...string) {
	a.set("segment", segments)
}

func (a *Audience) SetAbtest(abtests ...string) {
	a.set("abtest", abtests)
}

func (a *Audience) SetAlias(alias ...string) {
	a.set("alias", alias)
}

func (a *Audience) set(key string, v []string) {
	a.IsAll = false
	a.Value[key] = v
}
