// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"strings"
	"testing"

	"github.com/88250/lute"
)

var spinVditorIRBlockDOMTests = []*parseTest{

	{"15", "<p data-block=\"0\" data-node-id=\"20200813185628-tdpjvvr\">foo</p><p data-block=\"0\" data-node-id=\"20200813185636-sp1i1bp\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200813185628-tdpjvvr</span> <span class=\"vditor-ir__blockref\">\"bar\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> baz\n```\n[text](addr)\n```<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200813185628-tdpjvvr\">foo</p><p data-block=\"0\" data-node-id=\"20200813185636-sp1i1bp\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200813185628-tdpjvvr</span> <span class=\"vditor-ir__blockref\">&#34;bar&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> baz</p><div data-block=\"0\" data-node-id=\"20200813190155-d2yy4yv\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>[text](addr)<wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>[text](addr)</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"14", "<p data-block=\"0\" data-node-id=\"20200813113846-42h0ba1\">![foo](bar)<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200813113846-42h0ba1\"><span class=\"vditor-ir__node vditor-ir__node--expand\" data-type=\"img\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">foo</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><wbr><img src=\"bar\" alt=\"foo\"/></span></p>"},
	{"13", "<p data-block=\"0\" data-node-id=\"20200811153824-grm842y\">foo</p><blockquote data-block=\"0\" data-node-id=\"20200811153825-amjdbjz\"><p data-block=\"0\" data-node-id=\"\"><wbr><br></p></blockquote>", "<p data-block=\"0\" data-node-id=\"20200811153824-grm842y\">foo</p><p data-block=\"0\" data-node-id=\"20200811153825-amjdbjz\">&gt; <wbr></p>"},
	{"12", "<p data-block=\"0\" data-node-id=\"20200811153040-mrqu125\"><span data-type=\"a\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__link\">foo</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)b<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200811153040-mrqu125\"><span data-type=\"a\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__link\">foo</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span>b<wbr></p>"},
	{"11", "<h2 data-block=\"0\" class=\"vditor-ir__node vditor-ir__node--expand\" data-node-id=\"20200811142649-5osffhz\" id=\"ir-bar&lt;wbr&gt;\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\">## </span>foo<span data-type=\"heading-id\" class=\"vditor-ir__marker\"> {bar<wbr>}</span></h2>", "<h2 data-block=\"0\" class=\"vditor-ir__node vditor-ir__node--expand\" data-node-id=\"20200811142649-5osffhz\" id=\"ir-bar&lt;wbr&gt;\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\">## </span>foo<span data-type=\"heading-id\" class=\"vditor-ir__marker\"> {bar<wbr>}</span></h2>"},
	{"10", "<p data-block=\"0\" data-node-id=\"20200811140724-dxmm3jo\"><span data-type=\"em\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*</span><em data-newline=\"1\">foo</em><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*b<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200811140724-dxmm3jo\"><span data-type=\"em\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*</span><em data-newline=\"1\">foo</em><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*</span></span>b<wbr></p>"},
	{"9", "<p data-block=\"0\" data-node-id=\"20200810211034-9d34ae\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200810191413-6f5616</span> <span class=\"vditor-ir__blockref\">\"foo\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">) <wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200810211034-9d34ae\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200810191413-6f5616</span> <span class=\"vditor-ir__blockref\">&#34;foo&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> <wbr></p>"},
	{"8", "<p data-block=\"0\" data-node-id=\"20200809211933-b81ed7\">* foo<wbr></p>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200809211933-b81ed7\"><li data-marker=\"*\" data-node-id=\"\">foo<wbr></li></ul>"},
	{"7", "<p data-block=\"0\" data-node-id=\"20200809184752-a537de\">&gt; foo<wbr></p>", "<blockquote data-block=\"0\" data-node-id=\"20200809184752-a537de\"><p data-block=\"0\" data-node-id=\"\">foo<wbr></p></blockquote>"},
	{"6", "<p data-block=\"0\" data-node-id=\"20200809093825-b06abb\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">\"text<wbr>\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></p>", "<p data-block=\"0\" data-node-id=\"20200809093825-b06abb\"><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">&#34;text<wbr>&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></p>"},
	{"5", "<p data-block=\"0\" data-node-id=\"1596459249782\">foo ((bar)) <wbr></p>", "<p data-block=\"0\" data-node-id=\"1596459249782\">foo <span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span> <span class=\"vditor-ir__blockref\">&#34;bar&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> <wbr></p>"},
	{"4", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo \"text\")<wbr></p>\n", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo &#34;text&#34;)<wbr></p>"},
	{"3", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo \"text\"))<wbr></p>\n", "<p data-block=\"0\" data-node-id=\"1596459249782\"><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">&#34;text&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><wbr></p>"},
	{"2", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo))<wbr></p>\n", "<p data-block=\"0\" data-node-id=\"1596459249782\"><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">&#34;foo&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><wbr></p>"},
	{"1", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"2\"><wbr><br></p>", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"2\"><wbr></p>"},
	{"0", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"20200811112006-322210\"><wbr><br></p>", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"20200811112006-322210\"><wbr></p>"},
}

func TestSpinVditorIRBlockDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.BlockRef = true

	for _, test := range spinVditorIRBlockDOMTests {
		html := luteEngine.SpinVditorIRBlockDOM(test.from)

		if "15" == test.name {
			// 去掉最后一个新生成的节点 ID，因为这个节点 ID 是随机生成，每次测试用例运行都不一样，比较没有意义，长度一致即可
			lastNodeIDIdx := strings.LastIndex(html, "data-node-id=")
			length := len("data-node-id=\"20200813190226-k3m61yt\" ")
			html = html[:lastNodeIDIdx] + html[lastNodeIDIdx+length:]
			lastNodeIDIdx = strings.LastIndex(test.to, "data-node-id=")
			test.to = test.to[:lastNodeIDIdx] + test.to[lastNodeIDIdx+length:]
		}

		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
