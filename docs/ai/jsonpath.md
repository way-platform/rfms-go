# JSONPath -- XPath for JSON

## Copyright Notice

Copyright (c) 2020 IETF Trust and the persons identified as the document authors. All rights reserved.

This document is subject to BCP 78 and the IETF Trust's Legal Provisions Relating to IETF Documents ([https://trustee.ietf.org/license-info](https://trustee.ietf.org/license-info)) in effect on the date of publication of this document. Please review these documents carefully, as they describe your rights and restrictions with respect to this document. Code Components extracted from this document must include Simplified BSD License text as described in Section 4.e of the Trust Legal Provisions and are provided without warranty as described in the Simplified BSD License.

## [1.](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#section-1) [Introduction](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-introduction)

This document picks up the popular JSONPath specification dated 2007-02-21 [[JSONPath-orig](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#JSONPath-orig)] and provides a more normative definition for it. It is intended as a submission to the IETF DISPATCH WG, in order to find the right way to complete standardization of this specification. In its current state, it is a strawman document showing what needs to be covered.

### [1.1.](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#section-1.1) [Terminology](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-terminology)

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in BCP 14 [[RFC2119](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#RFC2119)] [[RFC8174](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#RFC8174)] when, and only when, they appear in all capitals, as shown here.

The grammatical rules in this document are to be interpreted as described in [[RFC5234](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#RFC5234)].The terminology of [[RFC8259](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#RFC8259)] applies.

Data Item:

A structure complying to the generic data model of JSON, i.e., composed of containers such as arrays and maps (JSON objects), and of atomic data such as null, true, false, numbers, and text strings.

Object:

Used in its generic sense, e.g., for programming language objects. When a JSON Object as defined in [[RFC8259](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#RFC8259)] is meant, we specifically say JSON Object.

Query:

Short name for JSONPath expression.

Argument:

Short name for the JSON data item a JSONPath expression is applied to.

Output Path:

A simple form of JSONPath expression that identifies a Position by providing a query that results in exactly that position. Similar to, but syntactically different from, a JSON Pointer [[RFC6901](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#RFC6901)].

Position:

A JSON data item identical to or nested within the JSON data item to which the query is applied to, expressed either by the value of that data item or by providing a JSONPath Output Path.

### [1.2.](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#section-1.2) [Inspired by XPath](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-inspired-by-xpath)

A frequently emphasized advantage of XML is the availability of plenty tools to analyse, transform and selectively extract data out of XML documents. [[XPath](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#XPath)] is one of these powerful tools.

In 2007, the need for something solving the same class of problems for the emerging JSON community became apparent, specifically for:- Finding data interactively and extracting them out of [[RFC8259](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#RFC8259)] data items without special scripting.
Specifying the relevant parts of the JSON data in a request by a client, so the server can reduce the data in the server response, minimizing bandwidth usage.So how does such a tool look like when done for JSON? When defining a JSONPath, how should expressions look like?The XPath expression

/store/book[1]/title

looks like

x.store.book[0].title

or

x['store']['book'][0]['title']

in popular programming languages such as JavaScript, Python and PHP, with a variable x holding the JSON data item. Here we observe that such languages already have a fundamentally XPath-like feature built in.The JSONPath tool in question should:- be naturally based on those language characteristics.
cover only essential parts of XPath 1.0.be lightweight in code size and memory consumption.be runtime efficient.

### [1.3.](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#section-1.3) [Overview of JSONPath Expressions](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-overview-of-jsonpath-expres)

JSONPath expressions always apply to a JSON data item in the same way as XPath expressions are used in combination with an XML document. Since a JSON data item is usually anonymous and doesn't necessarily have a "root member object", JSONPath used the abstract name `$` to refer to the top level object of the data item.

JSONPath expressions can use the _dot-notation_

$.store.book[0].title

or the _bracket-notation_

$['store']['book'][0]['title']

for paths input to a JSONPath processor. Where a JSONPath processor uses JSONPath expressions for internal purposes or as output paths, these will always be converted to the more general _bracket-notation_.JSONPath allows the wildcard symbol `*` for member names and array indices. It borrows the descendant operator `..` from [[E4X](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#E4X)] and the array slice syntax proposal `[start:end:step]` [[SLICE](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#SLICE)] from ECMASCRIPT 4.JSONPath can employ an _underlying scripting language_, expressions of which, written in parentheses: `(<expr>)`, can be used as an alternative to explicit names or indices as in:

$.store.book[(@.length-1)].title

The symbol `@` is used for the current object. Filter expressions are supported via the syntax `?(<boolean expr>)` as in

$.store.book[?(@.price < 10)].title

Here is a complete overview and a side by side comparison of the JSONPath syntax elements with its XPath counterparts.

[Table 1](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#table-1): [Overview over JSONPath, comparing to XPath](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-overview-over-jsonpath-comp)
|XPath|JSONPath|Description|
|---|---|---|
|`/`|`$`|the root object/element|
|`.`|`@`|the current object/element|
|`/`|`.` or `[]`|child operator|
|`..`|n/a|parent operator|
|`//`|`..`|nested descendants. JSONPath borrows this syntax from E4X.|
|`*`|`*`|wildcard. All objects/elements regardless of their names.|
|`@`|n/a|attribute access. JSON data items don't have attributes.|
|`[]`|`[]`|subscript operator. XPath uses it to iterate over element collections and for predicates. In JavaScript and JSON it is the native array operator.|
|`\|`|`[,]`|Union operator in XPath results in a combination of node sets. JSONPath allows alternate names or array indices as a set.|
|n/a|`[start:end:step]`|array slice operator borrowed from ES4.|
|`[]`|`?()`|applies a filter (script) expression.|
|n/a|`()`|script expression, using the underlying script engine.|
|`()`|n/a|grouping in Xpath|

XPath has a lot more to offer (location paths in unabbreviated syntax, operators and functions) than listed here. Moreover there is a significant difference how the subscript operator works in Xpath and JSONPath:

- Square brackets in XPath expressions always operate on the _node set_ resulting from the previous path fragment. Indices always start at 1.
With JSONPath, square brackets operate on the _object_ or _array_ addressed by the previous path fragment. Indices always start at 0.

## [2.](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#section-2) [JSONPath Examples](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-jsonpath-examples)

This section provides some more examples for JSONPath expressions. The examples are based on a simple JSON data item patterned after a typical XML example representing a bookstore (that also has bicycles):

```json
{
  "store": {
    "book": [
      { "category": "reference",
        "author": "Nigel Rees",
        "title": "Sayings of the Century",
        "price": 8.95
      },
      { "category": "fiction",
        "author": "Evelyn Waugh",
        "title": "Sword of Honour",
        "price": 12.99
      },
      { "category": "fiction",
        "author": "Herman Melville",
        "title": "Moby Dick",
        "isbn": "0-553-21311-3",
        "price": 8.99
      },
      { "category": "fiction",
        "author": "J. R. R. Tolkien",
        "title": "The Lord of the Rings",
        "isbn": "0-395-19395-8",
        "price": 22.99
      }
    ],
    "bicycle": {
      "color": "red",
      "price": 19.95
    }
  }
}
```

[Figure 1](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#figure-1): [Example JSON data item](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-example-json-data-item)

The examples in [Table 2](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#tbl-example) presume an underlying script language that allows obtaining the number of items in an array, testing for the presence of a map member, and performing numeric comparisons of map member values with a constant.

[Table 2](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#table-2): [Example JSONPath expressions applied to the example JSON data item](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-example-jsonpath-expression)
|XPath|JSONPath|Result|
|---|---|---|
|`/store/book/author`|`$.store.book[*].author`|the authors of all books in the store|
|`//author`|`$..author`|all authors|
|`/store/*`|`$.store.*`|all things in store, which are some books and a red bicycle.|
|`/store//price`|`$.store..price`|the price of everything in the store.|
|`//book[3]`|`$..book[2]`|the third book|
|`//book[last()]`|`$..book[(@.length-1)]`  <br>`$..book[-1:]`|the last book in order.|
|`//book[position()<3]`|`$..book[0,1]`  <br>`$..book[:2]`|the first two books|
|`//book[isbn]`|`$..book[?(@.isbn)]`|filter all books with isbn number|
|`//book[price<10]`|`$..book[?(@.price<10)]`|filter all books cheapier than 10|
|`//*`|`$..*`|all Elements in XML document. All members of JSON data item.|

## [3.](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#section-3) [Detailed definition](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-detailed-definition)

[TBD: This section needs to be fleshed out in detail. The text given here is intended to give the flavor of that detail, not to be the actual definition that is to be defined.]

JSONPath expressions, "queries" for short in this specification, are character strings, represented in UTF-8 unless otherwise required by the context in which they are used.When applied to a JSON data item, a query returns a (possibly empty) list of "positions" in the data item that match the JSONPath expression.

```
JSONPath = root *(step)
root = "$"

step = ".." name ; nested descendants
     / "." name ; child (dot notation)
     / "[" value-expression *("," value-expression) "]"
        ; child[ren] (bracket notation)
     / "[" value-expression *2(":" value-expression) "]"  ; (slice)
value-expression = *DIGIT / name
                 / script-expression / filter-expression
name = "'" text "'"
     / "*" ; wildcard
script-expression = "(" script ")"
filter-expression = "?(" script ")"
script = <To be defined in the course of standardization>
text = <any text, restrictions to be defined>
DIGIT = %x30-39
```

[Figure 2](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#figure-2): [ABNF definition for JSONPath](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-abnf-definition-for-jsonpat)

Within a script, @ stands for the position under consideration. [TBD: define underlying scripting language, if there is to be a standard one]

[TBD: define handling of spaces]A JSONPath starts at the root of the argument; the "current list" of positions is initialized to that root. Each step applies the semantics of that step to each of the positions in the "current list", returning another list; the "current list" is replaced by the concatenation of all these returned lists, and the next step begins. When all steps have been performed, the "current list" is returned, depending on the choices of the context either as a list of data items or as a list of output paths. [TBD: define the order of that list][TBD: Define all the steps][TBD: Define details of Output Path]

## [4.](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#section-4) [Discussion](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-discussion)

- Currently only single quotes allowed inside of JSONPath expressions.

Script expressions inside of JSONPath locations are currently not recursively evaluated by jsonPath. Only the global `$` and local `@` symbols are expanded by a simple regular expression.An alternative for jsonPath to return false in case of no match may be to return an empty array in future. [This is already done in the above.]

## [5.](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#section-5) [IANA considerations](https://www.ietf.org/archive/id/draft-goessner-dispatch-jsonpath-00.html#name-iana-considerations)

TBD: Define a media type for JSON Path expressions.
