# D4 Data Sorcerer

This is a fun little project to make gear upgrades easier to understand.  It's not a perfect system,
but it does provide a starting point.

The eventual goal is to empower ML to fine tune the weights and biases.  That's pretty far down the
road though.

## Initial Goals

- Intuitive enough to be used while playing.
- Simple enough for my wife to enjoy using.
- UI that is iPad Pro friendly.
- Robust enough to track the characters on the Eternal or Seasonal Realms.

### Stretch Goals

- CLI UI, because I'm a CLI junkie :heart:.

## The Method

The method is actually straightforward.  Each item slot has a ranking of stat priorities.
Each rank's weight is worth **0.10**.  The most important stat starts with **0.90** and each
subsequent rank is worth **0.10** less.  In order to factor in the **Item Power** we're adding
the __percentage of max__ to the sum of the weights.  Currently, the smart people of the internet
say that **~840** is the __max__ Item Power.  So we're adding `(itemPower / 840)`.

At the moment I'm not factoring in how well the stat rolled given the range, but you could factor
it in by using the `(rolledValue / maxValue)` instead of blindly using `1` for the value of the rank.

### Example

A Season 1 **Barbarian** following the **Walking Arsenal** guide will have the following stat
priorities for the **Helm** Slot.

| Rank | Stat               | Weight |
|------|--------------------|--------|
| 1    | Ranks of Upheaval  | 0.90   |
| 2    | Cooldown Reduction | 0.80   |
| 3    | % Total Armor      | 0.70   |
| 4    | Life on Kill       | 0.60   |
| 5    | Maximum Life       | 0.50   |
| 6    | Strength           | 0.40   |
| 7    | All Stats          | 0.30   |

We have an item with the following stats and the weights as described above:

| Colorful Clown Hat with Spinner  | Weight    |
|----------------------------------|-----------|
| 662+25 &rarr; **687** Item Power | 0.817     |
| **+3 Ranks** Upheaval            | 0.900     |
| 80 Dexterity                     | 0.000     |
| 10% Cooldown Reduction           | 0.800     | 
| 9.5% Total Armor                 | 0.70      |
| **TOTAL**                        | **3.217** |

So this example item will have a final value of **3.217**.
