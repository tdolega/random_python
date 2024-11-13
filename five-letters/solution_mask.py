def word_to_mask(word):
    mask = 0
    for letter in word:
        mask |= 1 << (ord(letter) - ord("a"))
    return mask


all_words = {}
with open("words_alpha.txt") as f:
    for word in f.read().split():
        if 5 == len(word) == len(set(word)):
            all_words[frozenset(word)] = word_to_mask(word)


def solve(letters_mask, words_masks):
    if letters_mask == 0:
        return 1

    letters_mask_count = 0
    rarest_mask = None
    min_count = float("inf")
    for i in range(26):
        letter_mask = 1 << i
        if letter_mask & letters_mask:
            letters_mask_count += 1
            count = 0
            for word_mask in words_masks:
                if word_mask & letter_mask:
                    count += 1

            if count < min_count:
                min_count = count
                rarest_mask = letter_mask

    n_solutions = 0
    for word_mask in words_masks + ([rarest_mask] if letters_mask_count % 5 else []):
        if word_mask & rarest_mask:
            n_solutions += solve(
                letters_mask=letters_mask & ~word_mask,
                words_masks=[wm for wm in words_masks if not wm & word_mask],
            )
    return n_solutions


answer = solve(
    letters_mask=(1 << 26) - 1,
    words_masks=list(all_words.values()),
)
print(answer)
