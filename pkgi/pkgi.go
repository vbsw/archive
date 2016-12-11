
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// Interfaces and functions to add package description.
package pkgi


type Version interface {
	Major() uint
	Minor() uint
	Patch() uint
	String() string
}

type About interface {
	Version() Version
	Authors() []string
	Name() string
	Imported() []string
	LicenseName() string
	LicenseContent() string
	Description() string
}

type Builder struct {
	version Version
	authors []string
	name string
	imported []string
	licenseName string
	licenseContent string
	description string
}

type tVersion struct {
	major uint
	minor uint
	patch uint
	str string
}

type tAbout struct {
	version Version
	authors []string
	name string
	imported []string
	licenseName string
	licenseContent string
	description string
}

func New() *Builder {
	return new(Builder)
}

func NewVersion(major, minor, patch uint) Version {
	version := new(tVersion)
	version.major = major
	version.minor = minor
	version.patch = patch
	version.str = VersionString(major, minor, patch)

	return version
}

func (this *Builder) Version(major, minor, patch uint) *Builder {
	this.version = NewVersion(major, minor, patch)

	return this
}

func (this *Builder) Authors(authors ...string) *Builder {
	this.authors = authors

	return this
}

func (this *Builder) Imported(imported ...string) *Builder {
	this.imported = imported

	return this
}

func (this *Builder) Name(name string) *Builder {
	this.name = name

	return this
}

func (this *Builder) License(licenseName, licenseContent string) *Builder {
	this.licenseName = licenseName
	this.licenseContent = licenseContent

	return this
}

func (this *Builder) Description(description string) *Builder {
	this.description = description

	return this
}

func (this *Builder) About() About {
	about := new(tAbout)
	about.version = this.version
	about.authors = this.authors
	about.name = this.name
	about.imported = this.imported
	about.licenseName = this.licenseName
	about.licenseContent = this.licenseContent
	about.description = this.description

	return about
}

func (this *tVersion) Major() uint {
	return this.major
}

func (this *tVersion) Minor() uint {
	return this.minor
}

func (this *tVersion) Patch() uint {
	return this.patch
}

func (this *tVersion) String() string {
	return this.str
}

func (this *tAbout) Version() Version {
	return this.version
}

func (this *tAbout) Authors() []string {
	return this.authors
}

func (this *tAbout) Name() string {
	return this.name
}

func (this *tAbout) Imported() []string {
	return this.imported
}

func (this *tAbout) LicenseName() string {
	return this.licenseName
}

func (this *tAbout) LicenseContent() string {
	return this.licenseContent
}

func (this *tAbout) Description() string {
	return this.description
}

func VersionString(major, minor, patch uint) string {
	if major < 10 && minor < 10 {

		if patch < 10 {
			versionArr := make([]byte, 5)
			versionArr[0] = byte(major + 48)
			versionArr[1] = 46
			versionArr[2] = byte(minor + 48)
			versionArr[3] = 46
			versionArr[4] = byte(patch + 48)

			return string(versionArr)

		} else if patch < 100 {
			versionArr := make([]byte, 6)
			patchTenth := patch / 10
			versionArr[0] = byte(major + 48)
			versionArr[1] = 46
			versionArr[2] = byte(minor + 48)
			versionArr[3] = 46
			versionArr[4] = byte(patchTenth + 48)
			versionArr[5] = byte(patch - patchTenth * 10 + 48)

			return string(versionArr)

		} else {
			majorByte := byte(major + 48)
			minorByte := byte(minor + 48)

			return string([]byte{majorByte, 46, minorByte, 46}) + uintToString(patch)
		}

	} else {
		return uintToString(major) + "." + uintToString(minor) + "." + uintToString(patch)
	}
}

func uintToString(value uint) string {
	var byteArr [20]byte
	var decimalPlaces int
	tenth := value / 10
	byteArr[19] = byte(value - tenth * 10 + 48)
	value = tenth

	for decimalPlaces = 1; value > 0 && decimalPlaces < 20; decimalPlaces += 1 {
		tenth := value / 10
		byteArr[19-decimalPlaces] = byte(value - tenth * 10 + 48)
		value = tenth
	}
	return string(byteArr[20-decimalPlaces:])
}
