package jujusvg

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gc "gopkg.in/check.v1"
	"gopkg.in/juju/charm.v6"
)

func Test(t *testing.T) { gc.TestingT(t) }

type newSuite struct{}

var _ = gc.Suite(&newSuite{})

var bundle = `
applications:
  mongodb:
    charm: "cs:precise/mongodb-21"
    num_units: 1
    annotations:
      "gui-x": "940.5"
      "gui-y": "388.7698359714502"
    constraints: "mem=2G cpu-cores=1"
  elasticsearch:
    charm: "cs:~charming-devs/precise/elasticsearch-2"
    num_units: 1
    annotations:
      "gui-x": "490.5"
      "gui-y": "369.7698359714502"
    constraints: "mem=2G cpu-cores=1"
  charmworld:
    charm: "cs:~juju-jitsu/precise/charmworld-58"
    num_units: 1
    expose: true
    annotations:
      "gui-x": "813.5"
      "gui-y": "112.23016402854975"
    options:
      charm_import_limit: -1
      source: "lp:~bac/charmworld/ingest-local-charms"
      revno: 511
relations:
  - - "charmworld:essearch"
    - "elasticsearch:essearch"
  - - "charmworld:database"
    - "mongodb:database"
series: precise
`

func iconURL(ref *charm.URL) string {
	return "http://0.1.2.3/" + ref.Path() + ".svg"
}

type emptyFetcher struct{}

func (f *emptyFetcher) FetchIcons(*charm.BundleData) (map[string][]byte, error) {
	return nil, nil
}

type errFetcher string

func (f *errFetcher) FetchIcons(*charm.BundleData) (map[string][]byte, error) {
	return nil, fmt.Errorf("%s", *f)
}

func (s *newSuite) TestNewFromBundle(c *gc.C) {
	b, err := charm.ReadBundleData(strings.NewReader(bundle))
	c.Assert(err, gc.IsNil)
	err = b.Verify(nil, nil)
	c.Assert(err, gc.IsNil)

	cvs, err := NewFromBundle(b, iconURL, nil)
	c.Assert(err, gc.IsNil)

	var buf bytes.Buffer
	cvs.Marshal(&buf)
	c.Logf("%s", buf.String())
	assertXMLEqual(c, buf.Bytes(), []byte(`
<?xml version="1.0"?>
<!-- Generated by SVGo -->
<svg width="631" height="457"
     style="font-family:Ubuntu, sans-serif;" viewBox="0 0 631 457"
     xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink">
<defs>
<g id="healthCircle" transform="scale(1.1)" >

<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16.000017 16.000017"><g transform="translate(-952 -156.362)"><path color="#000" overflow="visible" fill="none" d="M952 156.362h16v16h-16z"/><circle r="7.25" cy="164.362" cx="960" color="#000" overflow="visible" fill="#a7a7a7" stroke="#a7a7a7" stroke-width="1.5" stroke-dashoffset=".8"/><path style="line-height:125%;-inkscape-font-specification:Ubuntu;text-align:center" d="M963.8 161.286l-.066.057L959 165.49l-2.776-2.38-.84.948 3.616 3.804 5.5-5.787-.7-.79z" font-size="15" font-family="Ubuntu" letter-spacing="0" word-spacing="0" text-anchor="middle" fill="#fff"/></g></svg>
</g>
<svg:svg xmlns:svg="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" id="icon-1">
&#x9;&#x9;&#x9;&#x9;&#x9;<svg:image width="96" height="96" xlink:href="http://0.1.2.3/~juju-jitsu/precise/charmworld-58.svg"></svg:image>
&#x9;&#x9;&#x9;&#x9;</svg:svg><svg:svg xmlns:svg="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" id="icon-2">
&#x9;&#x9;&#x9;&#x9;&#x9;<svg:image width="96" height="96" xlink:href="http://0.1.2.3/~charming-devs/precise/elasticsearch-2.svg"></svg:image>
&#x9;&#x9;&#x9;&#x9;</svg:svg><svg:svg xmlns:svg="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" id="icon-3">
&#x9;&#x9;&#x9;&#x9;&#x9;<svg:image width="96" height="96" xlink:href="http://0.1.2.3/precise/mongodb-21.svg"></svg:image>
&#x9;&#x9;&#x9;&#x9;</svg:svg></defs>
<circle cx="47" cy="49" r="45" id="application-icon-mask" fill="none" />
<clipPath id="clip-mask" ><use x="0" y="0" xlink:href="#application-icon-mask" />
</clipPath>
<g id="relations">
<g >
<title>charmworld:essearch elasticsearch:essearch</title>
<line x1="413" y1="90" x2="90" y2="347" stroke="#a7a7a7" stroke-width="1px" stroke-dasharray="198.38, 16" />
<use x="243" y="210" xlink:href="#healthCircle" />
<circle cx="342" cy="146" r="4" fill="#a7a7a7" />
<circle cx="160" cy="290" r="4" fill="#a7a7a7" />
</g>
<g >
<title>charmworld:database mongodb:database</title>
<line x1="413" y1="90" x2="540" y2="366" stroke="#a7a7a7" stroke-width="1px" stroke-dasharray="143.91, 16" />
<use x="468" y="220" xlink:href="#healthCircle" />
<circle cx="450" cy="171" r="4" fill="#a7a7a7" />
<circle cx="502" cy="284" r="4" fill="#a7a7a7" />
</g>
</g>
<g id="applications">
<g transform="translate(323,0)" >
<title>charmworld</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<use x="0" y="0" xlink:href="#icon-1" transform="translate(42,42)" width="96" height="96" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >charmworld</text>
</g>
<g transform="translate(0,257)" >
<title>elasticsearch</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<use x="0" y="0" xlink:href="#icon-2" transform="translate(42,42)" width="96" height="96" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >elasticsearch</text>
</g>
<g transform="translate(450,276)" >
<title>mongodb</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<use x="0" y="0" xlink:href="#icon-3" transform="translate(42,42)" width="96" height="96" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >mongodb</text>
</g>
</g>
</svg>
`))
}

func (s *newSuite) TestNewFromBundleWithUnplacedApplication(c *gc.C) {
	b, err := charm.ReadBundleData(strings.NewReader(bundle))
	c.Assert(err, gc.IsNil)
	err = b.Verify(nil, nil)
	c.Assert(err, gc.IsNil)
	b.Applications["charmworld"].Annotations["gui-x"] = ""
	b.Applications["charmworld"].Annotations["gui-y"] = ""

	cvs, err := NewFromBundle(b, iconURL, nil)
	c.Assert(err, gc.IsNil)

	var buf bytes.Buffer
	cvs.Marshal(&buf)
	c.Logf("%s", buf.String())
	assertXMLEqual(c, buf.Bytes(), []byte(`
<?xml version="1.0"?>
<!-- Generated by SVGo -->
<svg width="901" height="290"
     style="font-family:Ubuntu, sans-serif;" viewBox="0 0 901 290"
     xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink">
<defs>
<g id="healthCircle" transform="scale(1.1)" >

<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16.000017 16.000017"><g transform="translate(-952 -156.362)"><path color="#000" overflow="visible" fill="none" d="M952 156.362h16v16h-16z"/><circle r="7.25" cy="164.362" cx="960" color="#000" overflow="visible" fill="#a7a7a7" stroke="#a7a7a7" stroke-width="1.5" stroke-dashoffset=".8"/><path style="line-height:125%;-inkscape-font-specification:Ubuntu;text-align:center" d="M963.8 161.286l-.066.057L959 165.49l-2.776-2.38-.84.948 3.616 3.804 5.5-5.787-.7-.79z" font-size="15" font-family="Ubuntu" letter-spacing="0" word-spacing="0" text-anchor="middle" fill="#fff"/></g></svg>
</g>
<svg:svg xmlns:svg="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" id="icon-1">
&#x9;&#x9;&#x9;&#x9;&#x9;<svg:image width="96" height="96" xlink:href="http://0.1.2.3/~juju-jitsu/precise/charmworld-58.svg"></svg:image>
&#x9;&#x9;&#x9;&#x9;</svg:svg><svg:svg xmlns:svg="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" id="icon-2">
&#x9;&#x9;&#x9;&#x9;&#x9;<svg:image width="96" height="96" xlink:href="http://0.1.2.3/~charming-devs/precise/elasticsearch-2.svg"></svg:image>
&#x9;&#x9;&#x9;&#x9;</svg:svg><svg:svg xmlns:svg="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" id="icon-3">
&#x9;&#x9;&#x9;&#x9;&#x9;<svg:image width="96" height="96" xlink:href="http://0.1.2.3/precise/mongodb-21.svg"></svg:image>
&#x9;&#x9;&#x9;&#x9;</svg:svg></defs>
<circle cx="47" cy="49" r="45" id="application-icon-mask" fill="none" />
<clipPath id="clip-mask" ><use x="0" y="0" xlink:href="#application-icon-mask" />
</clipPath>
<g id="relations">
<g >
<title>charmworld:essearch elasticsearch:essearch</title>
<line x1="810" y1="199" x2="90" y2="90" stroke="#a7a7a7" stroke-width="1px" stroke-dasharray="356.10, 16" />
<use x="442" y="136" xlink:href="#healthCircle" />
<circle cx="721" cy="185" r="4" fill="#a7a7a7" />
<circle cx="178" cy="103" r="4" fill="#a7a7a7" />
</g>
<g >
<title>charmworld:database mongodb:database</title>
<line x1="810" y1="199" x2="540" y2="109" stroke="#a7a7a7" stroke-width="1px" stroke-dasharray="134.30, 16" />
<use x="667" y="146" xlink:href="#healthCircle" />
<circle cx="724" cy="170" r="4" fill="#a7a7a7" />
<circle cx="625" cy="137" r="4" fill="#a7a7a7" />
</g>
</g>
<g id="applications">
<g transform="translate(720,109)" >
<title>charmworld</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<use x="0" y="0" xlink:href="#icon-1" transform="translate(42,42)" width="96" height="96" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >charmworld</text>
</g>
<g transform="translate(0,0)" >
<title>elasticsearch</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<use x="0" y="0" xlink:href="#icon-2" transform="translate(42,42)" width="96" height="96" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >elasticsearch</text>
</g>
<g transform="translate(450,19)" >
<title>mongodb</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<use x="0" y="0" xlink:href="#icon-3" transform="translate(42,42)" width="96" height="96" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >mongodb</text>
</g>
</g>
</svg>
`))
}

func (s *newSuite) TestWithFetcher(c *gc.C) {
	b, err := charm.ReadBundleData(strings.NewReader(bundle))
	c.Assert(err, gc.IsNil)
	err = b.Verify(nil, nil)
	c.Assert(err, gc.IsNil)

	cvs, err := NewFromBundle(b, iconURL, new(emptyFetcher))
	c.Assert(err, gc.IsNil)

	var buf bytes.Buffer
	cvs.Marshal(&buf)
	c.Logf("%s", buf.String())
	assertXMLEqual(c, buf.Bytes(), []byte(`
<?xml version="1.0"?>
<!-- Generated by SVGo -->
<svg width="631" height="457"
     style="font-family:Ubuntu, sans-serif;" viewBox="0 0 631 457"
     xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink">
<defs>
<g id="healthCircle" transform="scale(1.1)" >

<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16.000017 16.000017"><g transform="translate(-952 -156.362)"><path color="#000" overflow="visible" fill="none" d="M952 156.362h16v16h-16z"/><circle r="7.25" cy="164.362" cx="960" color="#000" overflow="visible" fill="#a7a7a7" stroke="#a7a7a7" stroke-width="1.5" stroke-dashoffset=".8"/><path style="line-height:125%;-inkscape-font-specification:Ubuntu;text-align:center" d="M963.8 161.286l-.066.057L959 165.49l-2.776-2.38-.84.948 3.616 3.804 5.5-5.787-.7-.79z" font-size="15" font-family="Ubuntu" letter-spacing="0" word-spacing="0" text-anchor="middle" fill="#fff"/></g></svg>
</g>
</defs>
<circle cx="47" cy="49" r="45" id="application-icon-mask" fill="none" />
<clipPath id="clip-mask" ><use x="0" y="0" xlink:href="#application-icon-mask" />
</clipPath>
<g id="relations">
<g >
<title>charmworld:essearch elasticsearch:essearch</title>
<line x1="413" y1="90" x2="90" y2="347" stroke="#a7a7a7" stroke-width="1px" stroke-dasharray="198.38, 16" />
<use x="243" y="210" xlink:href="#healthCircle" />
<circle cx="342" cy="146" r="4" fill="#a7a7a7" />
<circle cx="160" cy="290" r="4" fill="#a7a7a7" />
</g>
<g >
<title>charmworld:database mongodb:database</title>
<line x1="413" y1="90" x2="540" y2="366" stroke="#a7a7a7" stroke-width="1px" stroke-dasharray="143.91, 16" />
<use x="468" y="220" xlink:href="#healthCircle" />
<circle cx="450" cy="171" r="4" fill="#a7a7a7" />
<circle cx="502" cy="284" r="4" fill="#a7a7a7" />
</g>
</g>
<g id="applications">
<g transform="translate(323,0)" >
<title>charmworld</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<image x="42" y="42" width="96" height="96" xlink:href="http://0.1.2.3/~juju-jitsu/precise/charmworld-58.svg" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >charmworld</text>
</g>
<g transform="translate(0,257)" >
<title>elasticsearch</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<image x="42" y="42" width="96" height="96" xlink:href="http://0.1.2.3/~charming-devs/precise/elasticsearch-2.svg" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >elasticsearch</text>
</g>
<g transform="translate(450,276)" >
<title>mongodb</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<image x="42" y="42" width="96" height="96" xlink:href="http://0.1.2.3/precise/mongodb-21.svg" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >mongodb</text>
</g>
</g>
</svg>
`))
}

func (s *newSuite) TestDefaultHTTPFetcher(c *gc.C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<svg></svg>")
	}))
	defer ts.Close()

	tsIconUrl := func(ref *charm.URL) string {
		return ts.URL + "/" + ref.Path() + ".svg"
	}

	b, err := charm.ReadBundleData(strings.NewReader(bundle))
	c.Assert(err, gc.IsNil)
	err = b.Verify(nil, nil)
	c.Assert(err, gc.IsNil)

	cvs, err := NewFromBundle(b, tsIconUrl, &HTTPFetcher{IconURL: tsIconUrl})
	c.Assert(err, gc.IsNil)

	var buf bytes.Buffer
	cvs.Marshal(&buf)
	c.Logf("%s", buf.String())
	assertXMLEqual(c, buf.Bytes(), []byte(`
<?xml version="1.0"?>
<!-- Generated by SVGo -->
<svg width="631" height="457"
     style="font-family:Ubuntu, sans-serif;" viewBox="0 0 631 457"
     xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink">
<defs>
<g id="healthCircle" transform="scale(1.1)" >

<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16.000017 16.000017"><g transform="translate(-952 -156.362)"><path color="#000" overflow="visible" fill="none" d="M952 156.362h16v16h-16z"/><circle r="7.25" cy="164.362" cx="960" color="#000" overflow="visible" fill="#a7a7a7" stroke="#a7a7a7" stroke-width="1.5" stroke-dashoffset=".8"/><path style="line-height:125%;-inkscape-font-specification:Ubuntu;text-align:center" d="M963.8 161.286l-.066.057L959 165.49l-2.776-2.38-.84.948 3.616 3.804 5.5-5.787-.7-.79z" font-size="15" font-family="Ubuntu" letter-spacing="0" word-spacing="0" text-anchor="middle" fill="#fff"/></g></svg>
</g>
<svg:svg xmlns:svg="http://www.w3.org/2000/svg" id="icon-1"></svg:svg><svg:svg xmlns:svg="http://www.w3.org/2000/svg" id="icon-2"></svg:svg><svg:svg xmlns:svg="http://www.w3.org/2000/svg" id="icon-3"></svg:svg></defs>
<circle cx="47" cy="49" r="45" id="application-icon-mask" fill="none" />
<clipPath id="clip-mask" ><use x="0" y="0" xlink:href="#application-icon-mask" />
</clipPath>
<g id="relations">
<g >
<title>charmworld:essearch elasticsearch:essearch</title>
<line x1="413" y1="90" x2="90" y2="347" stroke="#a7a7a7" stroke-width="1px" stroke-dasharray="198.38, 16" />
<use x="243" y="210" xlink:href="#healthCircle" />
<circle cx="342" cy="146" r="4" fill="#a7a7a7" />
<circle cx="160" cy="290" r="4" fill="#a7a7a7" />
</g>
<g >
<title>charmworld:database mongodb:database</title>
<line x1="413" y1="90" x2="540" y2="366" stroke="#a7a7a7" stroke-width="1px" stroke-dasharray="143.91, 16" />
<use x="468" y="220" xlink:href="#healthCircle" />
<circle cx="450" cy="171" r="4" fill="#a7a7a7" />
<circle cx="502" cy="284" r="4" fill="#a7a7a7" />
</g>
</g>
<g id="applications">
<g transform="translate(323,0)" >
<title>charmworld</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<use x="0" y="0" xlink:href="#icon-1" transform="translate(42,42)" width="96" height="96" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >charmworld</text>
</g>
<g transform="translate(0,257)" >
<title>elasticsearch</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<use x="0" y="0" xlink:href="#icon-2" transform="translate(42,42)" width="96" height="96" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >elasticsearch</text>
</g>
<g transform="translate(450,276)" >
<title>mongodb</title>
<circle cx="90" cy="90" r="90" class="application-block" fill="#f5f5f5" stroke="#888" stroke-width="1" />
<use x="0" y="0" xlink:href="#icon-3" transform="translate(42,42)" width="96" height="96" clip-path="url(#clip-mask)" />
<rect x="0" y="135" width="180" height="32" rx="2" ry="2" fill="rgba(220, 220, 220, 0.8)" />
<text x="90" y="157" text-anchor="middle" style="font-weight:200" >mongodb</text>
</g>
</g>
</svg>
`))

}

func (s *newSuite) TestFetcherError(c *gc.C) {
	b, err := charm.ReadBundleData(strings.NewReader(bundle))
	c.Assert(err, gc.IsNil)
	err = b.Verify(nil, nil)
	c.Assert(err, gc.IsNil)

	ef := errFetcher("bad-wolf")
	_, err = NewFromBundle(b, iconURL, &ef)
	c.Assert(err, gc.ErrorMatches, "bad-wolf")
}

func (s *newSuite) TestWithBadBundle(c *gc.C) {
	b, err := charm.ReadBundleData(strings.NewReader(bundle))
	c.Assert(err, gc.IsNil)
	b.Relations[0][0] = "evil-unknown-application"
	cvs, err := NewFromBundle(b, iconURL, nil)
	c.Assert(err, gc.ErrorMatches, "cannot verify bundle: .*")
	c.Assert(cvs, gc.IsNil)
}

func (s *newSuite) TestWithBadPosition(c *gc.C) {
	b, err := charm.ReadBundleData(strings.NewReader(bundle))
	c.Assert(err, gc.IsNil)

	b.Applications["charmworld"].Annotations["gui-x"] = "bad"
	cvs, err := NewFromBundle(b, iconURL, nil)
	c.Assert(err, gc.ErrorMatches, `application "charmworld" does not have a valid position`)
	c.Assert(cvs, gc.IsNil)

	b, err = charm.ReadBundleData(strings.NewReader(bundle))
	c.Assert(err, gc.IsNil)

	b.Applications["charmworld"].Annotations["gui-y"] = "bad"
	cvs, err = NewFromBundle(b, iconURL, nil)
	c.Assert(err, gc.ErrorMatches, `application "charmworld" does not have a valid position`)
	c.Assert(cvs, gc.IsNil)
}
