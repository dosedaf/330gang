import streamlit as st
import matplotlib.pyplot as plt
import numpy as np

st.title("Multiple Charts in One Canvas")

# Generate data
x = np.linspace(0, 10, 100)
y1 = np.sin(x)
y2 = np.cos(x)

# Create a single figure with 2 subplots (2 columns)
fig, axes = plt.subplots(1, 2, figsize=(10, 4))

# First subplot (Sine wave)
axes[0].plot(x, y1, color='blue', label='Sine')
axes[0].set_title("Sine Wave")
axes[0].legend()

# Second subplot (Cosine wave)
axes[1].plot(x, y2, color='red', label='Cosine')
axes[1].set_title("Cosine Wave")
axes[1].legend()

# Display the Matplotlib figure in Streamlit
st.pyplot(fig)
