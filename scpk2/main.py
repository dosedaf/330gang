import pandas as pd
import streamlit as st
import sklearn as datasets 
import matplotlib.pyplot as plt

st.set_page_config(layout="wide")

if "df" not in st.session_state:
    st.session_state.df = pd.read_csv("community.csv")

df = st.session_state.df

st.title("Aplikasi Explorasi Dataset Communit")
st.write("menyajikan dataset iris secara interaktif")
length = st.sidebar.number_input("jumlah data", min_value=1, value=100, step=1)
# df = pd.read_csv("community.csv")
st.title("Dataset Community")
st.dataframe(df.head(length), width=1000000)

with st.expander("Informasi Dataset"):
    st.write("Dimensi Dataset : ", df.shape)
    st.write("Dimensi Statistik : ")
    st.write(df.describe())
    mis_val = df.isna().sum()
    st.write("Missing Values : ")
    st.write(mis_val)
    
continent_color = {
    "Asia": "Purple",
    "North America": "#D4ED26",
    "South America": "#1565C0",
    "Africa": "#64B5F6",
    "Australia": "Red",
    "Europe": "Green",
}
continents = df['Continent'].unique()
age, height = st.tabs(["Age", "Height"])

with age:
    st.header("Line Plot - Age per Index")
    cols = st.columns(2)
    for i, continent in enumerate(continents):
        if continent == "Antarctica":
            continue

        subset = df[df['Continent'] == continent]
    
        fig, ax = plt.subplots(figsize=(4, 2))
        ax.plot(subset.index, subset['Age'], marker='o', linestyle="-", color=continent_color[continent], label=continent)
        ax.set_xticks([0, 200, 400])
        ax.set_yticks([20, 40])
        
        ax.legend(loc="upper left")

        cols[i % 2].pyplot(fig)
with height:
    st.header("Line Plot - Height per Index")
    fig, ax = plt.subplots()
    
    for continent in continents:
        subset = df[df['Continent'] == continent]
        ax.plot(subset.index, subset["Height"], marker="o", linestyle="-", label=continent)
        
        ax.legend(title="Continent")
    
    st.pyplot(fig)