
//          Copyright 2017, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package window

type Properties struct {
	Width        int
	Height       int
	AntiAliasing bool
	Title        string
	Fullscreen   bool
}

func NewProperties() *Properties {
	properties := new(Properties)

	return properties
}

func (this *Properties) Clone() *Properties {
	properties := NewProperties()
	properties.Width = this.Width
	properties.Height = this.Height
	properties.AntiAliasing = this.AntiAliasing
	properties.Title = this.Title
	properties.Fullscreen = this.Fullscreen

	return properties
}
