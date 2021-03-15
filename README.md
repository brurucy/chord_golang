## Planning chord

### What we have done

* Node structure
* Ring Distance
* Closest hop

### Shortcuts

List of shortcuts <- contains all shortcuts, and during stabilize, if there's any nil pointers, it's just gonna drop them

### MigrateData

Needs to be run AFTER all shortcuts have been stabilized.

### Plan

1. Let's get the node logic down \\ Nikita


### What needs to be done

- [X] Fix Closest hop, take a better look at it, kinda messy, update it with the new correct distance function
- [ ] Fix AddShortcut, make it linear.
- [X] Fix findValue and find predecessor, it's probably fucked up
- [ ] Let's use a list of successor instead of Succ and SuccSucc pointers
- [ ] Stabilize, make sure that every time that the SUCC is updated, SUCCSUCC points to the SUCC of the SUCC, i.e if we add 7 between 5 and 17, then the succ of 5 is 7, and the succ succ is not gonna be 22, but 17 instead.
- [ ] Understand what it is and write tests for `MigrateData`
- [ ] Node logic:
  - [ ] `List`
  - [ ] `Lookup` Almost done. Left: 1) add count of # requests, 2) write tests
  - [ ] `Join`
  - [ ] `Leave`
  - [X] `Shortcut`
- [ ] **Networking**: DHT web server + Move node logic to threads/processes
- [ ] **CLI** (including input file parsing)
- [ ] BONUS: Make resilient to simulataneous leaves

### How to run go tests?

``` sh
cd src & go test
```
