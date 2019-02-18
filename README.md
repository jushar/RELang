# RELang
RELang is a domain-specific language to describe a memory interface to binary executables and transpiles to C++. It is based upon the ANTLR4 parser generator and is written in Golang.   
Its syntax looks very similar to C++.

## Motivation
Writing binary interfaces for existing applications you don't have the source of can be annoying with C++ and either needs a lot of boilerplate code or more generic approaches break auto-completion in IDEs.   
Therefore, it makes sense to create a small domain-specific language that lets you define binary interfaces and generates C++ code that can be easily embedded.

## Examples
```cpp
import "Element.relang";

class Vehicle : Element // Inherits from class 'Element'
{
    // Defines a function at the absolute address 0xBEAF
    void Fix() @ 0xBEAF;
    bool Attach(Element* element) @ 0x123456;

    // Defines an attribute at offset 4
    float32 health @ 0x4;

    // Based on the size and alignment of the previous attribute
    // this attribute is automatically defined to be at offset 8
    ModelInfo* model;

    // Line comment
    /* Block comment */
};
```
