# mctg

This project is a Markov chain text generator that uses Go for the core logic and Streamlit for the frontend. The generator learns from a provided text and generates new text based on the input.

## Features

- **Go Backend**: Utilizes Go for efficient text processing.
- **Python Integration**: Uses `ctypes` to interface with the Go shared library.
- **Streamlit Frontend**: Provides an easy-to-use web interface for interacting with the text generator.

## Requirements

- Python 3.6+
- Go 1.15+
- `streamlit`
- `ctypes`

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/kelbudiul/mctg.git
   cd markov-chain-text-generator
   ```

2. Install the required Python packages:

   ```bash
   pip install streamlit
   ```

3. Compile the Go code into a shared library:

   ```bash
   go build -o markov.dll -buildmode=c-shared markov.go
   ```

## Usage

1. Run the Streamlit app:

   ```bash
   streamlit run app.py
   ```

2. Open your browser and go to `http://localhost:8501`.

3. Enter training text in the provided text area, specify the start word and the length of the generated text, and click "Generate Text".

## Files

- `markov.go`: Contains the Go code for the Markov chain generator.
- `markov_chain.py`: Python interface to the Go shared library using `ctypes`.
- `app.py`: Streamlit app to provide a web interface for the text generator.
- `README.md`: Project documentation.

## Example

You can use the following example text to train the generator:

```text
Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do: once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, "and what is the use of a book," thought Alice, "without pictures or conversation?"

So she was considering in her own mind (as well as she could, for the hot day made her feel very sleepy and stupid), whether the pleasure of making a daisy-chain would be worth the trouble of getting up and picking the daisies, when suddenly a White Rabbit with pink eyes ran close by her.
```