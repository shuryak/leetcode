func nextGreatestLetter(letters []byte, target byte) byte {
    for _, val := range letters {
        if target < val {
            return val
        }
    }

    return letters[0]
}
