package main

// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN
// EDIT THE .haml FILE INSTEAD

import (
	"fmt"
	"html/template"
	"io"
	"os"
)

func NewEscapingWriter() (*EscapingWriter) {
	wr := &EscapingWriter{
		templates: make([]*template.Template, 0, len(escapingTemplatePatterns)),
	}
	
	for idx, pattern := range escapingTemplatePatterns {
		tmpl, err := template.New("escapingTemplates" + string(idx)).Parse(pattern)
		if err != nil {
			fmt.Errorf("Could not parse template: %d", idx)
			panic(err)
		}
		wr.templates = append(wr.templates, tmpl)
	}
	return wr
}

type EscapingWriter struct {
	templates []*template.Template
}

var escapingHtml = [...]string{
`<html>
	<body>
		<div></div>
		<div id="if">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
				`,
				`
				`,
				`
			</div>
		</div>
		<div id="else">
			<div class="exepected">
				`,
				`
			</div>
			<div></div>
			<div class="actual">
				`,
				`
				`,
				`
				`,
				`
			</div>
		</div>
		<div></div>
		<div id="rangeBody">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
				`,
				`
			</div>
		</div>
		<div id="constant">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="multipleAttrs">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlStartRel">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlStartAbsOk">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="protocolRelativeURLStart">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="pathRelativStart">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="dangerousURLStart">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="dangerousURLStart2">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="nonHierURL">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlPath">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlQuery">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlFragment">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsStrValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsNumericValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsBoolValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsNilValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsObjValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsObjValueScript">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsObjValueNotOverEscaped">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsStr">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="badMarshaler">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsMarshaler">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsStrNotUnderEscaped">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsRe">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsReBlank">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleBidiKeywordPassed">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleBidiPropNamePassed">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleExpressionBlocked">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleTagSelectorPassed">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleObfuscatedExpressionBlocked">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleObfuscatedMozBindingBlocked">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleURLQueryEncoded">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleURLBadProtocolBlocked">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="HtmlInText">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="HtmlInAttribute">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="HtmlInScript">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="HtmlInRCDATA">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="DynamicAttributeName">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
	</body>
</html>
`,
}

var escapingTemplatePatterns = []string{
	`Hello,`,
	`{{.C}}!`,
	`{{.H}}`,
	`{{.G}}`,
	`{{.}}`,
	`<a href="/search?q={{.Constant}}">`,
	`<a b=1 c={{.H}}>`,
	`<a href='{{.UrlStartRel}}'>`,
	`<a href='{{.UrlStartAbsOk}}'>`,
	`<a href='{{.ProtocolRelativeURLStart}}'>`,
	`<a href="{{.PathRelativeURLStart}}">`,
	`<a href='{{.DangerousURLStart}}'>`,
	`<a href=' {{.DangerousURLStart}}'>`,
	`<a href={{.NonHierURL}}>`,
	`<a href='http://{{.UrlPath}}/foo'>`,
	`<a href='/search?q={{.H}}'>`,
	`<a href='/faq#{{.H}}'>`,
	`<button onclick='alert({{.H}})'>`,
	`<button onclick='alert({{.N}})'>`,
	`<button onclick='alert({{.T}})'>`,
	`<button onclick='alert(typeof {{.Z}})'>`,
	`<button onclick='alert({{.A}})'>`,
	`<script>alert({{.A}})</script>"`,
	`<button onclick='alert({{.A}})'>`,
	`<button onclick='alert(&quot;{{.H}}&quot;)'>`,
	`<button onclick='alert(1/{{.B}}in numbers)'>`,
	`<button onclick='alert({{.M}})'>`,
	`<button onclick='alert({{.C}})'>`,
	`<button onclick='alert(/{{.JsRe}}/.test(""))'>`,
	`<script>alert(/{{.Blank}}/.test(""));</script>`,
	`<p style="dir: {{.Ltr}}">`,
	`<p style="border-{{.Left}}: 0; border-{{.Right}}: 1in">`,
	`<p style="width: {{"expression(alert(1337))"}}">`,
	`<style>{{.Selector}} { color: pink }</style>`,
	`<p style="width: {{.ObfuscatedExpression}}">`,
	`<p style="{{.ObfuscatedMozBinding}}: ...">`,
	`<p style="background: url(/img?name={{.Img}})">`,
	`<a style="background: url('{{.StyleURLBadProtocol}}')">`,
	`{{.W}}`,
	`<div title="{{.W}}">`,
	`<button onclick="alert({{.W}})">`,
	`<textarea>{{.W}}</textarea>`,
	`<input {{.Event}}="{{.Code}}">`,
}

func (wr *EscapingWriter) Execute(w io.Writer,
			data *TestDataType) {
	var err error = nil
		fmt.Fprint(w, escapingHtml[0])
	fmt.Fprint(w, "Hello, &lt;Cincinatti&gt;!")
	fmt.Fprint(w, escapingHtml[1])
	if data.T {
		fmt.Fprint(w, escapingHtml[2])
		err = wr.templates[0].Execute(w, data)
		handleEscapingError(err)
	}
	fmt.Fprint(w, escapingHtml[3])
	err = wr.templates[1].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[4])
	fmt.Fprint(w, "&lt;Goodbye&gt;")
	fmt.Fprint(w, escapingHtml[5])
	if data.F {
		fmt.Fprint(w, escapingHtml[6])
		err = wr.templates[2].Execute(w, data)
		handleEscapingError(err)
	} else {
		fmt.Fprint(w, escapingHtml[7])
		err = wr.templates[3].Execute(w, data)
		handleEscapingError(err)
	}
	fmt.Fprint(w, escapingHtml[8])
	fmt.Fprint(w, "&lt;a&gt;&lt;b&gt;")
	fmt.Fprint(w, escapingHtml[9])
	for _, d := range data.A {
		fmt.Fprint(w, escapingHtml[10])
		err = wr.templates[4].Execute(w, d)
		handleEscapingError(err)
	}
	fmt.Fprint(w, escapingHtml[11])
	fmt.Fprint(w, `<a href="/search?q=%27a%3cb%27">`)
	fmt.Fprint(w, escapingHtml[12])
	err = wr.templates[5].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[13])
	fmt.Fprint(w, "<a b=1 c=&lt;Hello&gt;>")
	fmt.Fprint(w, escapingHtml[14])
	err = wr.templates[6].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[15])
	fmt.Fprint(w, `<a href='/foo/bar?a=b&amp;c=d'>`)
	fmt.Fprint(w, escapingHtml[16])
	err = wr.templates[7].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[17])
	fmt.Fprint(w, `<a href='http://example.com/foo/bar?a=b&amp;c=d'>`)
	fmt.Fprint(w, escapingHtml[18])
	err = wr.templates[8].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[19])
	fmt.Fprint(w, `<a href='//example.com:8000/foo/bar?a=b&amp;c=d'>`)
	fmt.Fprint(w, escapingHtml[20])
	err = wr.templates[9].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[21])
	fmt.Fprint(w, `<a href="/javascript:80/foo/bar">`)
	fmt.Fprint(w, escapingHtml[22])
	err = wr.templates[10].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[23])
	fmt.Fprint(w, `<a href='#ZgotmplZ'>`)
	fmt.Fprint(w, escapingHtml[24])
	err = wr.templates[11].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[25])
	fmt.Fprint(w, `<a href=' #ZgotmplZ'>`)
	fmt.Fprint(w, escapingHtml[26])
	err = wr.templates[12].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[27])
	fmt.Fprint(w, `<a href=mailto:Muhammed%20%22The%20Greatest%22%20Ali%20%3cm.ali@example.com%3e>`)
	fmt.Fprint(w, escapingHtml[28])
	err = wr.templates[13].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[29])
	fmt.Fprint(w, `<a href='http://javascript:80/foo'>`)
	fmt.Fprint(w, escapingHtml[30])
	err = wr.templates[14].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[31])
	fmt.Fprint(w, `<a href='/search?q=%3cHello%3e'>`)
	fmt.Fprint(w, escapingHtml[32])
	err = wr.templates[15].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[33])
	fmt.Fprint(w, `<a href='/faq#%3cHello%3e'>`)
	fmt.Fprint(w, escapingHtml[34])
	err = wr.templates[16].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[35])
	fmt.Fprint(w, `<button onclick='alert(&#34;\u003cHello\u003e&#34;)'>`)
	fmt.Fprint(w, escapingHtml[36])
	err = wr.templates[17].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[37])
	fmt.Fprint(w, `<button onclick='alert( 42 )'>`)
	fmt.Fprint(w, escapingHtml[38])
	err = wr.templates[18].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[39])
	fmt.Fprint(w, `<button onclick='alert( true )'>`)
	fmt.Fprint(w, escapingHtml[40])
	err = wr.templates[19].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[41])
	fmt.Fprint(w, `<button onclick='alert(typeof null )'>`)
	fmt.Fprint(w, escapingHtml[42])
	err = wr.templates[20].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[43])
	fmt.Fprint(w, `<button onclick='alert([&#34;\u003ca\u003e&#34;,&#34;\u003cb\u003e&#34;])'>`)
	fmt.Fprint(w, escapingHtml[44])
	err = wr.templates[21].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[45])
	fmt.Fprint(w, `<script>alert(["\u003ca\u003e","\u003cb\u003e"])</script>`)
	fmt.Fprint(w, escapingHtml[46])
	err = wr.templates[22].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[47])
	fmt.Fprint(w, `<button onclick='alert([&#34;\u003ca\u003e&#34;,&#34;\u003cb\u003e&#34;])'>`)
	fmt.Fprint(w, escapingHtml[48])
	err = wr.templates[23].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[49])
	fmt.Fprint(w, `<button onclick='alert(&quot;\x3cHello\x3e&quot;)'>`)
	fmt.Fprint(w, escapingHtml[50])
	err = wr.templates[24].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[51])
	fmt.Fprint(w, `<button onclick='alert(1/ /* json: error calling MarshalJSON for type *template.badMarshaler: invalid character &#39;f&#39; looking for beginning of object key string */null in numbers)'>`)
	fmt.Fprint(w, escapingHtml[52])
	err = wr.templates[25].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[53])
	fmt.Fprint(w, `<button onclick='alert({&#34;\u003cfoo\u003e&#34;:&#34;O&#39;Reilly&#34;})'>`)
	fmt.Fprint(w, escapingHtml[54])
	err = wr.templates[26].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[55])
	fmt.Fprint(w, `<button onclick='alert(&#34;%3CCincinatti%3E&#34;)'>`)
	fmt.Fprint(w, escapingHtml[56])
	err = wr.templates[27].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[57])
	fmt.Fprint(w, `<button onclick='alert(/foo\x2bbar/.test(""))'>`)
	fmt.Fprint(w, escapingHtml[58])
	err = wr.templates[28].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[59])
	fmt.Fprint(w, `<script>alert(/(?:)/.test(""));</script>`)
	fmt.Fprint(w, escapingHtml[60])
	err = wr.templates[29].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[61])
	fmt.Fprint(w, `<p style="dir: ltr">`)
	fmt.Fprint(w, escapingHtml[62])
	err = wr.templates[30].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[63])
	fmt.Fprint(w, `<p style="border-left: 0; border-right: 1in">`)
	fmt.Fprint(w, escapingHtml[64])
	err = wr.templates[31].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[65])
	fmt.Fprint(w, `<p style="width: ZgotmplZ">`)
	fmt.Fprint(w, escapingHtml[66])
	err = wr.templates[32].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[67])
	fmt.Fprint(w, `<style>p { color: pink }</style>`)
	fmt.Fprint(w, escapingHtml[68])
	err = wr.templates[33].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[69])
	fmt.Fprint(w, `<p style="width: ZgotmplZ">`)
	fmt.Fprint(w, escapingHtml[70])
	err = wr.templates[34].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[71])
	fmt.Fprint(w, `<p style="ZgotmplZ: ...">`)
	fmt.Fprint(w, escapingHtml[72])
	err = wr.templates[35].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[73])
	fmt.Fprint(w, `<p style="background: url(/img?name=O%27Reilly%20Animal%281%29%3c2%3e.png)">`)
	fmt.Fprint(w, escapingHtml[74])
	err = wr.templates[36].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[75])
	fmt.Fprint(w, `<a style="background: url('#ZgotmplZ')">`)
	fmt.Fprint(w, escapingHtml[76])
	err = wr.templates[37].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[77])
	fmt.Fprint(w, `&iexcl;<b class="foo">Hello</b>, <textarea>O'World</textarea>!`)
	fmt.Fprint(w, escapingHtml[78])
	err = wr.templates[38].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[79])
	fmt.Fprint(w, `<div title="&iexcl;Hello, O&#39;World!">`)
	fmt.Fprint(w, escapingHtml[80])
	err = wr.templates[39].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[81])
	fmt.Fprint(w, `<button onclick="alert(&#34;&amp;iexcl;\u003cb class=\&#34;foo\&#34;\u003eHello\u003c/b\u003e, \u003ctextarea\u003eO&#39;World\u003c/textarea\u003e!&#34;)">`)
	fmt.Fprint(w, escapingHtml[82])
	err = wr.templates[40].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[83])
	fmt.Fprint(w, `<textarea>&iexcl;&lt;b class=&#34;foo&#34;&gt;Hello&lt;/b&gt;, &lt;textarea&gt;O&#39;World&lt;/textarea&gt;!</textarea>`)
	fmt.Fprint(w, escapingHtml[84])
	err = wr.templates[41].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[85])
	fmt.Fprint(w, `<input ZgotmplZ="doEvil()">`)
	fmt.Fprint(w, escapingHtml[86])
	err = wr.templates[42].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, escapingHtml[87])

	if err != nil {
		err = nil
	}
}

func handleEscapingError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
