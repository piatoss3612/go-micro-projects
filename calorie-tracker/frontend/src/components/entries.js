import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button, Form, Container, Modal } from "react-bootstrap";
import Entry from "./single-entry";

const emtyEntry = {
  dish: "",
  ingredients: "",
  calories: 0,
  fat: 0,
};

const Entries = () => {
  const [entries, setEntries] = useState([]);
  const [refreshData, setRefreshData] = useState(false);

  const [newEntry, setNewEntry] = useState(emtyEntry);
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
          setNewEntry(emtyEntry);
        }
      });
  };

  const deleteEntry = (id) => {
    var url = "http://localhost:8080/entry/delete/" + id;
    axios.delete(url, {}).then((response) => {
      if (response.status === 200) {
        setRefreshData(true);
      }
    });
  };

  const changeIngredientForEntry = () => {
    setChangeIngredient((prev) => ({ change: false, ...prev }));
    var url = "http://localhost:8080/entry/update/" + changeIngredient.id;
    axios.put(url, newEntry).then((response) => {
      if (response.status === 200) {
        setRefreshData(true);
        setNewEntry(emtyEntry);
      }
    });
  };

  const changeSingleEntry = () => {
    setChangeEntry((prev) => ({ change: false, ...prev }));
    var url = "http://localhost:8080/ingredient/update/" + changeEntry.id;
    axios
      .put(url, {
        ingredients: newIngredientName,
      })
      .then((response) => {
        if (response.status === 200) {
          setRefreshData(true);
          setNewIngredientName("");
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
      <Modal
        show={addNewEntry}
        onHide={() => {
          setAddNewEntry(false);
        }}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Add Calorie Entry</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Dish</Form.Label>
            <Form.Control
              onChange={(evt) =>
                setNewEntry((prev) => ({ dish: evt.target.value, ...prev }))
              }
            />
            <Form.Label>Ingredients</Form.Label>
            <Form.Control
              onChange={(evt) =>
                setNewEntry((prev) => ({
                  ingredients: evt.target.value,
                  ...prev,
                }))
              }
            />
            <Form.Label>Calories</Form.Label>
            <Form.Control
              onChange={(evt) =>
                setNewEntry((prev) => ({ calories: evt.target.value, ...prev }))
              }
            />
            <Form.Label>Fat</Form.Label>
            <Form.Control
              type="number"
              onChange={(evt) =>
                setNewEntry((prev) => ({ fat: evt.target.value, ...prev }))
              }
            />
          </Form.Group>
          <Button className="" onClick={() => addEntry()}>
            Add
          </Button>
          <Button onClick={() => setAddNewEntry(false)}>Cancel</Button>
        </Modal.Body>
      </Modal>
      <Modal
        show={changeIngredient.change}
        onHide={() => setChangeIngredient({ change: false, id: 0 })}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Change Ingredients</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>New Ingredients</Form.Label>
            <Form.Control
              onChange={(evt) => setNewIngredientName(evt.target.value)}
            />
          </Form.Group>
          <Button onClick={changeIngredientForEntry}>Change</Button>
          <Button
            onClick={() =>
              setChangeIngredient((prev) => ({ change: false, ...prev }))
            }
          >
            Cancel
          </Button>
        </Modal.Body>
      </Modal>
      <Modal
        show={changeEntry.change}
        onHide={() => setChangeEntry({ change: false, id: 0 })}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Change Entry</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Dish</Form.Label>
            <Form.Control
              onChange={(evt) =>
                setNewEntry((prev) => ({ dish: evt.target.value, ...prev }))
              }
            />
            <Form.Label>Ingredients</Form.Label>
            <Form.Control
              onChange={(evt) =>
                setNewEntry((prev) => ({
                  ingredients: evt.target.value,
                  ...prev,
                }))
              }
            />
            <Form.Label>Calories</Form.Label>
            <Form.Control
              onChange={(evt) =>
                setNewEntry((prev) => ({ calories: evt.target.value, ...prev }))
              }
            />
            <Form.Label>Fat</Form.Label>
            <Form.Control
              type="number"
              onChange={(evt) =>
                setNewEntry((prev) => ({ fat: evt.target.value, ...prev }))
              }
            />
          </Form.Group>
          <Button onClick={changeSingleEntry}>Change</Button>
          <Button
            onClick={() =>
              setChangeEntry((prev) => ({ change: false, ...prev }))
            }
          >
            Cancel
          </Button>
        </Modal.Body>
      </Modal>
    </div>
  );
};

export default Entries;
