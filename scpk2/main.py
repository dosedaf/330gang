import pandas as pd
import streamlit as st
import matplotlib.pyplot as plt

# biar lebih lebar content area nya
st.set_page_config(layout="wide")

# biar ga trigger reload kalo ada hubungannya sama df (tabs)
if "df" not in st.session_state:
    st.session_state.df = pd.read_csv("community.csv")
df = st.session_state.df

st.title("Aplikasi Explorasi Dataset Community")
st.write("Menyajikan dataset iris secara interaktif")
length = st.sidebar.number_input("jumlah data", min_value=1, value=100, step=1)
st.title("Dataset Community")
st.dataframe(df.head(length), width=1000000) # limit data ditampilkan

with st.expander("Informasi Dataset"):
    st.write("Dimensi Dataset : ", df.shape) # bisa concatenate ternyata
    st.write("Dimensi Statistik : ")
    st.write(df.describe())
    mis_val = df.isna().sum()
    st.write("Missing Values : ")
    st.write(mis_val)
    
# set warna tiap benua
continent_color = {
    "Asia": "Purple",
    "North America": "#D4ED26",
    "South America": "#1565C0",
    "Africa": "#64B5F6",
    "Australia": "Red",
    "Europe": "Green",
}

# dapetin array of continents
continents = df['Continent'].unique()
age, height = st.tabs(["Age", "Height"]) # buat tabs

with age:
    st.subheader("Line Plot - Age per Index")

    # buang antarctica
    age_continents = [c for c in continents if c != "Antarctica"]
    
    l = len(age_continents)
    cols = 2 
    rows = (l + 1) // cols

    fig, axes = plt.subplots(rows, cols, figsize=(8, 2 * rows), sharex=True, sharey=True)
    axes = axes.flatten() # digepengin rows dan cols nya

    for i, continent in enumerate(age_continents):
        subset = df[df['Continent'] == continent]

        ax = axes[i]  # ambil current ax
        ax.plot(subset.index, subset['Age'], marker='o', linestyle="-", 
                color=continent_color.get(continent, "black"), label=continent)
        ax.set_title(continent)
        ax.legend(loc="upper left")
        ax.set_xticks([0, 200, 400])
        ax.set_yticks([20, 40])

    st.pyplot(fig)

with height:
    st.subheader("Line Plot - Height per Index")
    fig, ax = plt.subplots(figsize=(11, 6))
    
    for continent in continents:
        subset = df[df['Continent'] == continent]
        ax.plot(subset.index, subset["Height"], marker="o", linestyle="-", label=continent)
        
        ax.legend(title="Continent")
    
    st.pyplot(fig)