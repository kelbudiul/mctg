import streamlit as st
from markov_chain import MarkovChain

# Title of the app
st.title("Markov Chain Text Generator")

# Text input for training data
training_data = st.text_area("Enter training text", height=200)

# Input for start word and length of generated text
start_word = st.text_input("Enter the start word")
length = st.number_input("Enter the length of generated text", min_value=1, max_value=100, value=10)

# Button to train the model and generate text
if st.button("Generate Text"):
    # Create a MarkovChain instance
    mc = MarkovChain()
    
    # Add training data to the chain
    mc.add(training_data)
    
    # Generate text
    generated_text = mc.generate(start_word, length)
    
    # Display generated text
    st.subheader("Generated Text")
    st.write(generated_text)
