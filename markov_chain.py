import ctypes
import os

# Load the shared library into ctypes
lib = ctypes.CDLL(os.path.abspath("markov.dll"))

# Define the MarkovChain class to interact with the Go library
class MarkovChain:
    def __init__(self):
        lib.NewMarkovChain.restype = ctypes.c_int
        self.obj_id = lib.NewMarkovChain()

    def add(self, text):
        lib.Add.argtypes = [ctypes.c_int, ctypes.c_char_p]
        lib.Add(self.obj_id, ctypes.c_char_p(text.encode('utf-8')))

    def generate(self, start, length):
        lib.Generate.argtypes = [ctypes.c_int, ctypes.c_char_p, ctypes.c_int]
        lib.Generate.restype = ctypes.c_char_p
        result = lib.Generate(self.obj_id, ctypes.c_char_p(start.encode('utf-8')), length)
        return result.decode('utf-8')
