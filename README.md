# Markov Chain Text Generator

This project is a Markov chain-based text generator that combines the efficiency of Go for core logic with the user-friendliness of Python and Streamlit for the frontend. The generator learns from provided text and generates new content based on the learned patterns.

## Features

- **Go Backend**: Utilizes Go for efficient text processing and generation.
- **Python Integration**: Uses `ctypes` to interface with the Go shared library.
- **Streamlit Frontend**: Provides an intuitive web interface for interacting with the text generator.
- **Cross-language Implementation**: Demonstrates effective integration between Go and Python.

## Requirements

- Docker Engine
- Python 3.7+
- Go 1.15+
- Streamlit
- ctypes (included in Python standard library)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/kelbudiul/mctg.git
   cd mctg
   ```

2. Build the Docker image:

   ```bash
   docker build -t mctg .
   ```

## Usage

1. Run the Docker container:

   ```bash
   docker run --name mctg-client -p 8501:8501 mctg
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

## How It Works

1. The Go code implements the core Markov chain functionality.
2. The Go code is compiled into a shared library (.dll file).
3. Python uses ctypes to load and interact with the Go shared library.
4. Streamlit provides a web interface for users to input text and generate new content.
5. The Markov chain learns from the input text and generates new text based on the learned probabilities.

## Example

You can use the following sample text to test the generator:

```text
Once upon a time, in the heart of a vast kingdom, there lived a wise old king who ruled over the land with fairness and kindness. The people adored their king, for under his reign, peace and prosperity had blossomed. Every morning, the sun would rise over the hills, casting a golden hue upon the villages and farms, and the fields would burst forth with crops.

But as the days went by, a dark shadow began to creep over the land. Whispers of unrest stirred in the farthest reaches of the kingdom, where the wind carried tales of a mysterious figure lurking in the forests. Some said it was a sorcerer, banished long ago for practicing dark arts. Others believed it to be nothing more than a myth, a story spun by fearful minds. Nevertheless, the rumors spread, and with them, unease grew.

The king, knowing that his people needed reassurance, called upon his most trusted knights to investigate. Among them was Sir Reginald, a brave and noble warrior who had fought many battles in service of the crown. With his sword at his side and a determined heart, Sir Reginald ventured into the forest, where the air grew thick with silence and the trees seemed to whisper secrets of their own.

As he ventured deeper, the path became less clear, twisting and turning as though the forest itself were alive. Sir Reginald paused to catch his breath, only to hear a faint sound, like the rustling of leaves. He turned, his hand resting on the hilt of his sword, and saw a figure emerge from the shadows. It was a woman, cloaked in black, her eyes glimmering with an unnatural light.

"You seek the truth," she said in a voice that seemed to echo through the trees. "But the truth is not always what it seems. What you find here may change the fate of the kingdom forever."

Sir Reginald, though wary, stepped forward. "Who are you? And what do you know of the darkness that plagues our land?"

The woman smiled, a slow and knowing smile. "I am but a messenger. The answers you seek lie beyond, in the heart of the forest. But beware, for the choices you make here will have consequences."

With that, she vanished, leaving Sir Reginald alone once more. He pressed on, driven by the weight of his duty and the promise of answers. The forest grew denser, the air colder, until finally, he reached a clearing. In the center stood an ancient stone, covered in runes that glowed faintly in the moonlight.

As Sir Reginald approached, a voice echoed in his mind: "The balance of the world is fragile. Power, when unchecked, can bring ruin. Choose wisely."

He reached out, his hand hovering over the stone. In that moment, he realized that his choice would not only determine his fate but the fate of the entire kingdom.
```

## Acknowledgments

- This project demonstrates the effective integration of Go and Python in a text generation application.
- Thanks to the Streamlit team for providing an excellent tool for creating user interfaces with Python.
