all_words = {}
with open("words_alpha.txt", "r") as f:
    for word in f:
        word = word.strip()
        if 5 == len(word) == len(set(word)):
            all_words[frozenset(word)] = word


def solve(letters, words, solution):
    if not letters:
        return 1
    words_concat = "".join(words)
    least_common_letter = min(letters, key=words_concat.count)
    n_solutions = 0
    fill = [least_common_letter] if len(letters) % 5 else []
    for word in words + fill:
        if least_common_letter in word:
            sword = set(word)
            n_solutions += solve(
                letters=letters - sword,
                words=[word for word in words if sword.isdisjoint(word)],
                solution=solution + [word],
            )
    return n_solutions


alphabet = set((chr(i) for i in range(ord("a"), ord("z") + 1)))
answer = solve(alphabet, list(all_words.values()), [])
print(answer)
