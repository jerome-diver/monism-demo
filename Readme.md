# Monism demonstration

## Definition

A Monism here is an interface contract for

- a variadic function as Exec func(...Monism) Monism
- a static function as ID() string
  and a function to modify any printable Monism with other specific ones.

It can have different kind of Monism's type to exist together to create something.

Monism identity is dependent of his Exec action and ID return statically a name reflecting his identity to be used by.

This Exec action can:

- print a content of a printable Monism
- transform an existing Monism into
- transform the meta print of the printable Monism

The ID is just a static string to return his identity as we want top define for.

Now with function that get Monism and return Monism, you would be able to build something to be Idiomatic and easy to use.

## Design Pattern to use for and to use with

Monism use many design:

- Variadic function
- Pure function
- High Order Function
- Closure
- Composit structure and function
- Decorator design pattern
- Strategy design pattern
  All of these achieve to be used with Functional code to produce with Option design.

## Advantages

- Easy to test
- Easy to read and understand
- Flexible
- Idiomatic
- Maintainable the easy way
