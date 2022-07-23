import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button, Form, Container, Modal } from "react-bootstrap";
import Entry from "./single-entry";

const Entries = () => {
  const [entries, setEntries] = useState([]);
  const [refreshData, setRefreshData] = useState(false);

  const [newEntry, setNewEntry] = useState({
    dish: "",
    ingredients: "",
    calories: 0,
    fat: 0,
  });
  const [addNewEntry, setAddNewEntry] = useState(false);
  const [changeEntry, setChangeEntry] = useState({ change: false, id: 0 });

  const [newIngredientName, setNewIngredientName] = useState("");
  const [changeIngredient, setChangeIngredient] = useState({
    change: false,
    id: 0,
  });

  const getAllEntries = () => {
    var url = "http://localhost:8080/entries";
    axios
      .get(url, {
        responseType: "json",
      })
      .then((response) => {
        if (response.status === 200) {
          setEntries(response.data);
        }
      });
  };

  useEffect(() => {
    getAllEntries();
  }, []);

  if (refreshData) {
    setRefreshData(false);
    getAllEntries();
  }

  const addEntry = () => {
    setAddNewEntry(false);
    var url = "http://localhost:8080/entry/create";
    axios
      .post(url, {
        ingredients: newEntry.ingredients,
        dish: newEntry.dish,
        calories: newEntry.calories,
        fat: parseFloat(newEntry.fat),
      })
      .then((response) => {
        if (response.status === 200) {
          setRefreshData(true);
        }
      });
  };

  const deleteEntry = (id) => {
    var url = "http://localhost:8080/entry/delete" + id;
    axios.delete(url, {}).then((response) => {
      if (response.status === 200) {
        setRefreshData(true);
      }
    });
  };

  return (
    <div>
      <Container>
        <Button onClick={() => setAddNewEntry(true)}>
          Track Today's Calories
        </Button>
      </Container>
      <Container>
        {entries != null &&
          entries.map((entry, i) => (
            <Entry
              key={i}
              entryData={entry}
              deleteEntry={deleteEntry}
              setChangeEntry={setChangeEntry}
              setChangeIngredient={setChangeIngredient}
            ></Entry>
          ))}
      </Container>
    </div>
  );
};

export default Entries;
