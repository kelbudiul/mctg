# Markov Chain Text Generator

This project is a Markov chain-based text generator that combines the efficiency of Go for core logic with the user-friendliness of Python and Streamlit for the frontend. The generator learns from provided text and generates new content based on the learned patterns.

## Features

- **Go Backend**: Utilizes Go for efficient text processing and generation.
- **Python Integration**: Uses `ctypes` to interface with the Go shared library.
- **Streamlit Frontend**: Provides an intuitive web interface for interacting with the text generator.
- **Cross-language Implementation**: Demonstrates effective integration between Go and Python.

## Requirements

- Python 3.6+
- Go 1.15+
- Streamlit
- ctypes (included in Python standard library)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/markov-chain-text-generator.git
   cd markov-chain-text-generator
   ```

2. Install the required Python package:
   ```bash
   pip install streamlit
   ```

3. Compile the Go code into a shared library:
   ```bash
   go build -o markov.dll -buildmode=c-shared markov.go
   ```

## Usage

1. Start the Streamlit app:
   ```bash
   streamlit run app.py
   ```

2. Open your web browser and navigate to `http://localhost:8501`.

3. In the web interface:
   - Enter your training text in the provided text area.
   - Specify a start word for the generated text.
   - Set the desired length of the generated text.
   - Click "Generate Text" to create new text based on your input.

## Project Structure

- `markov.go`: Go implementation of the Markov chain logic.
- `markov.h`: C header file for the Go shared library.
- `markov_chain.py`: Python wrapper for interacting with the Go library.
- `app.py`: Streamlit application providing the user interface.
- `README.md`: This file, containing project documentation.

## How It Works

1. The Go code implements the core Markov chain functionality.
2. The Go code is compiled into a shared library (.dll file).
3. Python uses ctypes to load and interact with the Go shared library.
4. Streamlit provides a web interface for users to input text and generate new content.
5. The Markov chain learns from the input text and generates new text based on the learned probabilities.

## Example

You can use the following sample text to test the generator:

```
Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do: once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, "and what is the use of a book," thought Alice, "without pictures or conversation?"

So she was considering in her own mind (as well as she could, for the hot day made her feel very sleepy and stupid), whether the pleasure of making a daisy-chain would be worth the trouble of getting up and picking the daisies, when suddenly a White Rabbit with pink eyes ran close by her.
```

## Acknowledgments

- This project demonstrates the effective integration of Go and Python in a text generation application.
- Thanks to the Streamlit team for providing an excellent tool for creating user interfaces with Python.