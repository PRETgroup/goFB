# An introduction to Textual Function Blocks

Textual Function Blocks allow for simple and quick development of FBs in IEC 61499.

This document will briefly explain how the language works.

## Introduction to IEC61499 and TFB

IEC61499 is the newest Industrial Automation standard for Programmable Logic Controllers. It is not a programming language, rather, it is a specification, like HTML. Just as HTML describes how a website should appear and what it should comtain, IEC61499 describes how PLCs can be instructed to behave.

Programming in IEC61499 consists of describing Function Blocks (FBs). FBs are individual units of behaviour and state, much like a `class` is in an object-oriented programming language such as C++. 

In C++, we could define a class `dog`. The `dog` might have methods to call on it, such as `feed()` and `walk()`, and then we could invoke those methods from an external controller. 
In IEC61499, we could define a Function Block `dog`, and give it input events `feed` and `walk`.

In short, Function Blocks are designed to encapsulate models, allowing for the creation of reusuable components. 
There are three main kinds:
* _Basic Function Blocks_, which contain a single Execution Control Chart (state machine) which describes how the FB should behave, given some inputs, internal variables, and state.
* _Service Interface Function Blocks_, which are implementation-specific blocks allowing for interfacing with underlying hardware and/or environmental signals.
* _Composite Function Blocks_, which contain an internal network of other Basic and Composite Function Blocks.

Between these two types of block, very complex systems can be realised.

In the dog example, in C++, we might define a class `dog`, and a class `house` that contains an instance of the class `dog`. Our top level file would be some `main.cpp` which instantiated a `house` and then ran the program.

In IEC61499, we can define a basic FB `dog`, and a composite FB `house` that contains an instance of the basic FB `dog`. Our top level file could simply be an instantiation of `house`, or some other top level composite FB that contained `house`, which then runs the program.

### Why use IEC61499?

In industrial automation, it is undesirable to have errors in programmed systems, as these could lead to catastrophic consequences when considering the plants and processes involved. While it is of course possible to write industrial controllers in languages like C++, it can be more difficult to validate and verify designs. 

IEC61499 has had plenty of work done to streamline design and design validation for industrial automation processes. It's a good tool for the job.

### Why use Textual Function Blocks for IEC61499?

IEC61499 was primarily designed to be programmed in a visual environment, and many GUI/WSIWYG tools for IEC61499 have been developed. This in and of itself this is not a bad thing, but there are three main drawbacks to visual approaches:
1. When rendering complex systems in WYSIWYGs, the visual representations naturally tend to become complex, encapsulating lots of minute details. These rapidly can become difficult to parse and understand.
2. GUI-based tools tend to be more complex, less platform-independent, more intensive, and less interoperable than command-line based alternatives.
3. Visual representations tend to have complex filetypes, which may not play nicely with version control software. Text-based representations of FBs allows for simpler and clearer integrations.

Just as Structured Text (ST) is a valid alternative for representing Ladder Logic and/or Function Block control systems in the IEC 61131-3 standard, Textual Function Blocks seek to become a valid alternative for those working with the newer IEC 61499 standard.

## Setting up your workspace

TFB files are text files with the \*.tfb extension.

Your project should be based out of a single TFB file or folder containing TFB files. 
A single TFB file can contain one or many function blocks.

## General syntax

* In TFBs, all commands are terminated with semicolons. 
* The language is not whitespace sensitive. 
* Multiple commands can be put on a single line. 
* Curly braces denote scope.
* The language requires forward declaration of FBs, that is, FB names must be declared before the block is defined.
* The language is case-sensitive.
* Comments begin with //
* There are no block comments

## Block declaration

To begin with, your block will be declared. This will look like the following, for an IEC61499 Basic Function Block called firstBasic:

```
basicFB firstBasic;
```

Or, for an IEC61499 Composite Function Block called firstComposite:

```
compositeFB firstComposite;
```

Or, for an IEC61499 Service Interface Function Block called firstService:

```
serviceFB firstService;
```

## Declaring the interface

Interfaces and architectures of blocks are separated in TFBs. An interface declaration is a scoped list of input and output events and data for function blocks.

Composite and BasicFBs have nearly identical declaration syntax. 

An example, a composite FB with input event `a`, output event `b`, input integer `c`, and output integer `d`:

```
compositeFB firstComposite;
interface of firstComposite {
    in event a;
    out event b;
    in int c;
    out int d;
}
```

### Associations

The only difference in the interfaces for basic, service, and composite FBs is that basicFBs/serviceFBs can have *event-data associations*, whereas compositeFBs can't.

An example of event data associations on a basicFB:
```
basicFB firstBasic;
interface of firstBasic {
    in event a;
    out event b;
    in int c with a;
    out int d with c;
}
```

Event-data associations are used on basicFBs in IEC61499 to identify which data lines should be updated when certain events occur.

Every name field in the interface supports comma-separated values for more concise code. Hence, the following two examples, where there are two event inputs `a` and `b`, are equivalent.

```
basicFB firstBasic1;
interface of firstBasic1 {
    in event a;
    in event b;
}

basicFB firstBasic2;
interface of firstBasic2 {
    in event a,b;
}
```

For instance, the following will create four input events `a`, `b`, `c`, `d`, and six inputs `num1`, `num2`, `num3`, `num4`, `num5`, `num6`. 

The associations work on a many-to-many relationship, so `num2` is associated to both `b` and `c`. 

Likewise, `num5` and `num6` are both associated to both `a` and `b`.

```
basicFB firstBasic;
interface of firstBasic {
    in event a, b, c, d;
    in int num1 with a;
    in int num2 with b, c;
    in int num3, num4 with d;
    in int num5, num6, with a, b;
}
```

### Valid Types

Types are taken directly from the IEC61499 specification. For more information on the types, refer to the spec.

Valid event types are
* event

Valid data types are 
* bool
* byte
* word
* dword
* lword
* sint
* usint
* int
* uint
* dint
* udint
* lint
* ulint
* real
* lreal
* time

### Arrays and Default/Initial Values

Array sizes can be specified on any valid data type. For instance, the following will create an input integer `y` with an array size of 2, and associate it with input event `a`.

```
basicFB firstBasic;
interface of firstBasic {
    in event a;
    in int[2] y with a;
}
```

Initial values can also be specified on any valid data type. For instance, the following will set the initial value on `z` to be `0`. 

```
basicFB firstBasic;
interface of firstBasic {
    in event a;
    in int initial 0 z with a;
}
```
When initial values are not provided, they will default to some implementation-specific value.

If multiple values are presented on that line, all will be given that default value. For instance, this sets the initial values on both `z` and `y` to be `0`.

```
basicFB firstBasic;
interface of firstBasic {
    in event a;
    in int initial 0 y,z with a;
}
```

Initial values can also be provided for arrays. These should be surrounded by square brackets `[` and `]`.

```
basicFB firstBasic;
interface of firstBasic {
    in event a;
    in int[3] initial [0,1,2] z with a;
}
```

Array sizes and initial conditions can also be used on compositeFBs.

### The HouseDog so far

In our introduction, we described a `dog` in a `house`. We can implement that in TFBs. Here's what we might have so far.

```
basicFB dog;
interface of dog {
    in event tick; //we call this to update the dog's behaviour periodically

    in event feed;
    in int initial 0 foodweight_g with feed; //the amount the dog will eat, in grams
    in event walk;
    in int initial 0 walkDistance_m with walk; //the distance the dog will walk, in metres

    out event bark; //emit this when the dog barks
   
    out event weightChange;
    out int initial 7500 dogweight_g with weightChange; //the dogs weight in grams
}
```

## Basic Function Block Architectures

Baic Function Block architectures in TFBs consist of internal variables (denoted `internals`), states in a state machine (denoted `states`), and named algorithms (denoted `algorithms`).

States can contain invokations of named and anonymous algorithms, as well as output event emission, and output transitions to other states.

We'll now go through this in more detail.

An architecture is a scoped list of these things, just like an interface. Here is an empty example (not that this won't compile due to the elipses):

_For the purposes of remaining succinct, assume all blocks `firstBasic` are defined with the interface in this example._

```
basicFB firstBasic;
interface of firstBasic {
    in event a;
    out event b;
    in int x with a;
    out int y with b;
}
architecture of firstBasic {
    internals {
        ...
    }
    states {
        ...
    }
    algorithms {
        ...
    }
}
```

### Internal variables

Internals have the same syntax as interface data lines, except that they have no direction. For instance, if we have an internal integer variables `p` and `q` initial `0`, and internal variable boolean type `t`:

```
...
architecture of firstBasic {
    internals {
        int initial 0 p,q;
        bool t;
    }
}
```

We can express these individually as well, using `internal` instead of `internals`. The following code compiles identically to the previous:

```
...
architecture of firstBasic {
    internal int initial 0 p,q;
    internal bool t;
}
```

### Execution Control Chart (State Machine) States and Transitions

States represent the different nodes in the execution control chart (state machine) inside the basicFB. Each state is named, and can have transitions to travel between states. For instance, the following state machine would switch between states `s1` and `s2` every time an `a` was received, and it would emit `b` each time it entered `s2`.
```
...
architecture of firstBasic {
    states {
        s1 {
            -> s2 on a;
        }
        s2 {
            emit b;
            -> s1 on a;
        }
    }
}
```
Just like with internals, we can express states individually by using `state` instead of `states`. The following code compiles identically to the previous:
```
...
architecture of firstBasic {
    state s1 {
        -> s2 on a;
    }
    
    state s2 {
        emit b;
        -> s1 on a;
    }
}
```

The initial state in the state machine is simply the one that appears first.

Transitions are denoted by the `->` operand. This points to the state they travel to. They have their condition after the `on` keyword, and the condition must be terminated by a semicolon. If no condition is provided, the default condition, `true`, is used.

Conditions can be as complex as needed, and support referencing internal and internal variables, and input events. For instance, in the following code, the transition from `s1` to `s2` is only when `a` has occured, and also when input variable `x` is greater than or equal to 5 and less than 15.

```
...
architecture of firstBasic {
    state s1 {
        -> s2 on a && x >= 5 && x < 15;
    }
    
    state s2 {
        emit b;
        -> s1 on a;
    }
}
```

In IEC61499, states perform their outputs only on entry to the state. In the previous code, `s2` will emit `b` only when it is entered. It will not continue to emit `b` forever.

In the state, all actions are performed in the order that they appear. Then, all the transitions are evaluated, in the order that they appear.

The `emit` statement, just like many other statements in TFBs, supports chaining comma-separated names (e.g. `emit a,b,c;`). 

The `->` transition operand does not support chaining, as this would not make sense.

### Named Algorithms

We can also link arbitrary code to states. In IEC61499, this is done through the usage of named algorithms. 

IEC61499 does not specify what language algorithms are written in, so here we use C. Presented in the following example is an algorithm called `increment_y` which increments the output variable `y` by the value of the input variable `x`. It is called each time `s2` is entered.

```
...
architecture of firstBasic {
    states {
        s1 {
            -> s2 on a;
        }
        s2 {
            emit b;
            run increment_y;
            -> s1 on a;
        }
    }

    algorithms {
        increment_y in "C" `me->y += me->x`;
    }
}
```

*(It is important to note that the compiler used for your IEC61499 generated code needs to support your algorithm language choice. In our case, goFB supports C in algorithms, and variables are all referenced using the `me->[var_name]` notation).*

As can be seen in the previous example, algorithms have a specified language, as well as a name, and algorithm contents (surrounded by backticks). 

The `run` statement, just like many other statements in TFBs, supports chaining comma-separated names (e.g. `run increment_y, another_algorithm;`).

Just as with `states` and `state` and `internals` and `internal`, algorithms can also be individually specified by using `algorithm` instead of `algorithms`. In the above example, this would be

```
    ...
    algorithm increment_y in "C" `me->y += me->x`;
}
```

### Anonymous algorithms

Often in IEC61499, algorithms can be short and used only in a single state. To this end, TFBs support anonymous algorithms, allowing you to specify short algorithms as needed. In the `increment_y` example, we can replace the named algorithm with an anonymous inline algorithm to make the code more concise:

```
...
architecture of firstBasic {
    states {
        s1 {
            -> s2 on a;
        }
        s2 {
            emit b;
            run in "C" `me->y += me->x`;
            -> s1 on a;
        }
    }
}
```

### The HouseDog so far

We can combine all of the above to create the state machine for our dog. In this, our dog has a default wait state `st_wait`.

From `st_wait`, one of three things can happen:
1. The dog is fed, so the dog gains weight equal to the food eaten.
2. The dog is walked, so the dog loses weight equal to the distance travelled in metres. The dog can't go below 5kg in weight however.
3. The dog has nothing happen for 10 ticks, so he barks, and resets the counter.

```
basicFB dog;
interface of dog {
    in event tick; //we call this to update the dog's behaviour periodically

    in event feed;
    in int initial 0 foodweight_g with feed; //the amount the dog will eat, in grams
    in event walk;
    in int initial 0 walkDistance_m with walk; //the distance the dog will walk, in metres

    out event bark; //emit this when the dog barks
   
    out event weightChange;
    out int initial 7500 dogweight_g with weightChange; //the dogs weight in grams
}

architecture of dog {
    internal int initial 0 tickCount;

    states {
        st_wait {
            run incTick;

            -> st_feed on feed;
            -> st_walk on walk;
            -> st_bark on tickCount == 10;
            -> st_wait on tick;
        }

        st_feed {
            run in "C" `me->dogweight_g += me->foodweight_g`;
            run clrTick;
            emit weightChange;

            -> st_wait;
        }

        st_walk {
            run walkDog;
            run clrTick;
            emit weightChange;

            -> st_wait;
        }

        st_bark {
            run clrTick;
            emit bark;
            -> st_wait;
        }
    }

    algorithms {
        clrTick in "C" `me->tickCount = 0;`;
        incTick in "C" `me->tickCount++;`;
    } 
        
    algorithm walkDog in "C" `
        me->dogweight_g -= me->walkDistance_m; 
        if(me->dogweight_g < 5000) {
            me->dogweight_g = 5000;
        }
    `;
        
}
```

## Service Interface Function Block Architectures

Service Interface Function Blocks (SIFBs) in IEC61499 consist only of an interface declaration. 

They are usually provided with implementation files for the compilation process, or reference some built-in libraries in the compiler. They are usually implementation specific (i.e. dependent on a given compiler/target platform).

Here is a simple SIFB. In this example, we would need to be providing an additional file `SIFB_counter.h` (to be imported by the other files as needed). We might also have a `SIFB_counter.c` which provided the implementation of the SIFB to gcc/clang.
When writing the \*.c and \*.h files, refer to the documentation of your compiler.

```
serviceFB counter compileheader "SIFB_counter.h";
interface of counter {
	out event countup;
	out int count_value with countup; 
}
```

A special extension for goTFB is that it also supports autogeneration of certain SIFBs when using the goFB IEC61499 compiler.

Here is an empty example of a SIFB in the TFB language for the SIFB format that goFB uses:

```
serviceFB counter;
interface of counter {
	out event countup;
	out int count_value with countup;
}

architecture of counter {
	in "C";

	in_struct ``;
	pre_init ``;
	init ``;
	run ``;
	shutdown ``;
}
```

This architecture specification is rigidly defined. The `in` keyword function must be the first element of the architecture, and applies to all raw code locations `in_struct`, `pre_init`, `init`, `run`, and `shutdown`. 

These locations refer to different areas in the execution lifestyle compiled by goFB. They are, as follows:
* `in_struct` - Placed inside the structure that defines the serviceFB. Use this to store variable declarations, etc.
* `pre_init` - First initialisation pass. Default values present in the interface will already be loaded on ports, however, no communication from other FBs or loading of external data has occured yet. Use this for startup code not dependent on other modules being initialised.
* `init` - Second initialisation pass. Configuration data possibly provided from other FBs has now been loaded on ports.
* `run` - This runs every update of the overall FB system. Use this for managing I/O and creating events during runtime.
* `shutdown` - This runs when the FB system is shutting down. May not be necessary for "always online" embedded systems.

An example of an implemented counter, which simply emits a new count every update:

```
serviceFB counter;
interface of counter {
	out event countup;
	out int initial 0 count_value with countup;
}

architecture of counter {
	lang "C";

	run `me->count_value++; me->outputEvents.event.countup = 1;`;
}
```

## Composite Function Block Architectures

Composite Function Block architectures in TFBs consist of internal FB instances (denoted `instances`), connections between instance event ports (denoted `events`), and connections between instance data ports (denoted `data`).

Data connections can also be connected to constant parameters.

We'll now go through this in more detail.

An architecture is a scoped list of these things, just like an interface. Here is an empty example (note that this won't compile due to the elipses):

_For the purposes of remaining succinct, assume all blocks `firstBasic` and `firstComposite` are defined with the interfaces in this example._

```
basicFB firstBasic;
interface of firstBasic {
    in event a;
    out event b;
    in int x with a;
    out int y with b;
}

...

compositeFB firstComposite;
interface of firstComposite {
	in event c;
	out event d;
	in int p; 
	out int q;
}

architecture of firstComposite {
	instances {
		firstBasic myFirst1;
		firstBasic myFirst2;
	}
	events {
		myFirst1.a <- c;
		myFirst2.a <- myFirst1.b;
		d <- myFirst2.b;
	}
	data {
		myFirst1.x <- p;
		myFirst2.x <- myFirst1.y;
		q <- myFirst2.y;
	}

}
```

TODO the rest of this section