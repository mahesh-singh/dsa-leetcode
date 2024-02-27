class WordDictionary:
    def __init__(self):
            self.root = {}
    def search(self, word: str) -> bool:
        if word == "me":
            return True
        return False

if __name__ == "__main__":
    obj = WordDictionary()
    word = "abc"
    m = {}
    m['a'] = "a"

    print(m['a'])
    print(m.get('a') != None)
