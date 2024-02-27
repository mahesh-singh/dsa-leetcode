class TreeNode:
    def __init__(self):
        self.children = {}
        self.is_word = False


class WordDictionary:

    def __init__(self):
        self.root = TreeNode()

    def addWord(self, word: str) -> None:
        node = self.root
        for char in word:
            if char not in node.children:
                #set the value
                node.children.setdefault(char, TreeNode())
            #assign to next 
            node = node.children[char]
        node.is_word = True

    def search(self, word: str) -> bool:

        def dfs(node, index):
            if len(word) == index:
                return node.is_word
            if word[index] == ".":
                #Getting the child node only not the key and value both from map
                for child in node.children.values():
                    if dfs(child, index+1):
                        return True
            #search char in node.children's key from map
            if word[index] in node.children:
                return dfs(node.children[word[index]], index+1)
            return False
        
        return dfs(self.root, 0)
