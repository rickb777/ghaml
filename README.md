Ghaml - the golang parser for a HAML-like language
=====

Ghaml is a go parser for a haml-like language. It's goals are to provide a
HAML-like syntax for creating views while being fast and efficient.

Ghaml achieves this by compiling ghaml templates into go code, which can then
be compiled into your application directly. To ease this burden, the ghaml tool
can call 'go build' for you, making a near seemless replacement.

[![Build Status](https://travis-ci.org/travissimon/ghaml.png)](https://travis-ci.org/travissimon/ghaml)

What is HAML?
=============
Haml stands for HTML Abstraction Markup Language, a templating system originally
developed for the Ruby language. Ghaml is very similar in intent, differing where
it makes sense to accomodate the differences between Go and Ruby.

HAML's core principles are:
1 Markup Should be Beautiful
2 Markup Should be DRY
3 Markup Should be Well-Indented
4 HTML Structure Should be Clear

Learn more about Haml here: http://haml.info/

A look at a Template
====================

    @data_type: string

    %html
      %head
        %title= "Hello, ", data
      %body
        %h1= "Hello, ", data
        
        %div
          This is child content for the div above. Note that
		  HAML is space-sensitive, so all text indented at
          this level is encased in the div.
    
        #id_div
          You can use the # operator as a shortcut to create
          a div with the given id.

        .implicit_class
          The .operator (think of the '.' css selector') lets
          you create a div with the given class. For example
          this text will be wrapped in a div that looks like
          this: `<div class="implicit_class">...`
		
        %ul{type:disc}
          - for i := 0; i < 10; i++ { // arbitrary go code
            %li= "Item: ", i
          - }

The ghaml template above illustrates many features of Ghaml templates:

* Strongly-typed data type: the @data_type directive specifies that the 
  template accepts a data parameter of type `string`. This will typically
  be a struct of view content.
* Outputting tag content via the `=` operator is syntactically equal to
  the variadic parameter definition of the `fmt.Print()` function. Therefore
  variables and strings can be concatenated by seperating them with commas.
  Further, the '=' operator protects against Cross-Site Scripting (XSS) attacks.
* The `-` operator lets the developer execute arbitrary Go code.
* The special '|' operator is equal to the '=' operator, but does not perform
  HTML encoding, allowing for formatted output, with the risk of XSS attacks.

Ghaml command syntax
====================

The default workflow is to use the ghaml command in your working directory.
This will compile all ghaml template, and then run the 'go build' command.
Command line options are:

* -v Output verbose (as in 'some') comments on what is going on
* -clean Remove all generated *.go files
* -nogo Do not run the 'go build' command. This might be useful in a
   build script
* -f forces template regeneration


Copyright / License
===================
Copyright 2013 Travis Simon

Licensed under the GNU General Public License Version 2.0 (or later); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.gnu.org/licenses/old-licenses/gpl-2.0.txt

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.