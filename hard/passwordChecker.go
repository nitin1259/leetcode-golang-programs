package hard

import "fmt"

/**
A password is considered strong if below conditions are all met:

It has at least 6 characters and at most 20 characters.
It must contain at least one lowercase letter, at least one uppercase letter, and at least one digit.
It must NOT contain three repeating characters in a row ("...aaa..." is weak, but "...aa...a..." is strong, assuming other conditions are met).
Write a function strongPasswordChecker(s), that takes a string s as input, and return the MINIMUM change required to make s a strong password.
If s is already strong, return 0.

Insertion, deletion or replace of any one character are all considered as one change.
*/

func strongPasswordChecker(s string) int {
	fmt.Println("Strong Password checker")

	l := len(s)

	if l < 6 {
		return 6 - l
	} else if l > 20 {
		return l - 20
	}

	return 0
}

/**

Approach to the problem:

Hi everyone,

Instead of providing my crap code, I thought it would be more helpful to discuss how I approached this problem.

Note: I'm assuming that everyone is able to detect what password constraints are satisfied / violated using an O(n) pass of the proposed password. I won't explain how to do that. Instead I'll explain how to use that information obtained to answer the question.

First off, terminology. I will refer to our three password constraints by these letters:

Constraint A: Length
Constraint B: Uppercase/Lowercase/Digit
Constraint C: Repeating characters

I viewed this problem as having three distinct cases based on where we are with respect to Constraint A (password length). Here they are in order of complexity:

Character count is between 6 and 20. All we need to do is figure out how many edits we need.
Character count is less than 6. We need to raise it to 6 and maybe edit a few characters depending on other constraints.
Character count is greater than 20. We need to remove characters until we get down to 20 and maybe edit a few characters depending on other constraints.
Before we discuss these cases individually, I want to have a brief note about Constraint C. Specifically, I want to show that for breaking up repeating sequences, editing a character is more efficient than adding or removing characters. Quick visual proof:

"aaaaaaaa" --> 8 characters
"aa1aa1aa" --> Solved with 2 edits. In general we need (L/3) edits for a sequence length of L.
"aa1aa1aa1aa" --> Solved with 3 additions. In general we need (L/2 - 1) additionss for a sequence length of L.
"aa" --> Solved with 6 removes. In general we need (L - 2) removes for a sequence length of L.

Now that we got that out of the way, let's discuss the three cases individually:

===== Case 1 =====

In this case, Constraint A is already satisfied so we don't need to do any adds or removes, only edits. To minimize the number of edits, we want each edit to ideally work toward satisfying both Constraint B and Constraint C, which is more than possible. We can replace a character in the middle of a repeating sequence with a character from a missing character group (uppercase, lowercase, digit). Two birds, one stone! Whichever constraint needs more edits to resolve is going to be our answer here.

For Constraint B, it's relatively simple to figure out how many character edits you need to have all character groups.

Note: Don't worry about the edits you make for Constrint B violating Constraint C. In a password <20 characters, you will always have enough character options to edit in without introducing repeats, so this just isn't an issue.
For Constraint C, recall that the most efficient solution for a sequence length of L is L/3 edits, as discussed above.

Note: Don't worry about the edits you make for Constrint C violating Constraint B. You are changing one of multiple repeating characters, so there are other instances of that character group. For example, changing "aaa" to "a2a" does not remove all "a"s, so it can't possibly violate Constraint B by removing all lowercase characters.
Given that you need X characters edited to satisfy constraint B, and Y characters edited to satisfy constraint C, and they have the two-birds-one-stone relationship described above, the answer is Max(X,Y).

===== Case 2 =====

In this case, we need to add a certain amount of characters (at most 6) to satisfy Constraint A. While we're doing that, we may as well add specific characters to satify Constraint B, two birds with one stone! In fact, this "adding" can also break up repeating character sequences, so three birds one stone! Example:

"aaaGG" --> Needs 1 more character for constraint A, needs digit for constraint B, needs to break up "aaa" for Constraint C.
"aa1aGG" --> All three constraints satisfied with a single edit. :)

Given this three-bird-one-stone approach, just calculating how many characters you need for Constraint A, Constraint B and Constraint C individually and taking the Max() of them is enough to solve this subproblem.

Note that for Constraint C, you need to use (L/2 - 1) additions instead of (L/3) edits. Adjust your formula accordingly!
===== Case 3 =====

In this case, we need to remove a certain amount of characters (until password length is down to 20) to satisfy Constraint A. Unfortunately Constraint A and Constraint B are completely independent (one requires explicit removes, the other requires edits), but they both have independent two-birds-one-stone relationships with Constraint C (for which both edits and removes are potential solutions). This makes Constraint C very complicated to accurately calculate. We'll get to this below.

First off, we can calculate the number of removes needed for Constraint A and the number of edits needed for Constraint B. We already did similar work for Cases 1 & 2.

Now, for Constraint C, we need to recall that in general, editing is a more efficient way to break up repeating sequences than removing. However, sometimes a remove can be as valuable as an edit. Example:

"aaaaaccceeennn"
"aac1ce1en1n" --> [Bad] Used 3 removes, 3 edits. Specifically, removed three "a"s, then broke up the other letters with edits.
"aa1aacceenn" --> [Good] Used 3 removes, 1 edit. Specifically, removed one "c", "e" and "n", and then broke up the "a"s with a single edit.

See how we used the same amount of removes, but by being smart about it, we saved up to 2 edits? This is important, because we already have to remove a set number of characters to satisfy Constraint A. We may as well two-birds-one-stone them to reduce the number of edits we will need for Constraint C.

To see the relationship between the number of edits vs removes needed to fix a repeating sequence, look below:

"aaaaaaaaaa" --> Repeat length 10.
"aa1aa1aa1a" --> Used 3 edits, 0 removes to satisfy constraint.
"aa1aa1aa" --> Used 2 edits, 2 removes to satisfy constraint.
"aa1aa" --> Used 1 edit, 5 removes to satisfy constraint.
"aa" --> Used 0 edits, 8 removes to satisfy constraint.

Observation 1: In general, 1 edit is worth 3 removes as expected.
Observation 2: For the solution with 3 edits, the 3rd edit only saved us 2 removes instead of 3.
Observation 3: The exact number of removes the last edit will save us is determined by sequence length.
Observation 4: We never insert/edit/remove the first two characters of the sequence. We only care about the remaining L-2 characters.
If you look at the example and he subsequent observations, you will eventually figure out that the last edit to break up a sequence of length L is worth (L-2)%3 removes. This is because, again, we only care about L-2 characters, and each previous edit broke up 3 of the L-2 characters.

Given these realizations, and given the number of characters you have to remove for Constraint A, here's the logic to use them as efficiently as possible to reduce the number of edits needed for Constraint C:

For each sequence where (L-2 % 3) == 1, one remove saves you an edit. Remove 1 character from these sequences first!
E.g. for "bbb", L=3, L-2 % 3 = 1, you need 1 remove to not need an edit)
For each sequence where (L-2 % 3) == 2, two removes save you an edit. Remove 2 characters from these sequences next!
E.g. for "bbbb", L=4, L-2 % 3 = 2, you need 2 removes to not need an edit)
For each sequence where (L-2 % 3) == 0, three removes save you an edit. Remove 3 characters at a time from these sequences.
E.g. for "bbbbb", L=5, L-2 % 3 = 0, you need 3 removes to not need an edit)
Now that you removed as much as you needed to remove for Constraint A, recalculate the repeating sequence lengths (L). For each new L, you will need L/3 edits, as we already know from Cases 1 & 2.

By now we established that you are going to:

Remove X characters for Constraint A
Then edit Y characters for Constraint C (given that we already removed the most efficient X characters from the repeating sequences)
We also know we need to edit Z characters for Sequence B, but Y and Z can now be two-birds-one-stone'd with a Max() function since we calculated Y completely independently from Constraint A. So the final answer is (X + Max(Y,Z)).
*/
