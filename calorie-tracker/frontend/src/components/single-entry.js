import React from "react";
import "bootstrap/dist/css/bootstrap.css";
import { Button, Card, Row, Col } from "react-bootstrap";

const Entry = ({
  entryData,
  setChangeEntry,
  setChangeIngredient,
  deleteEntry,
}) => {
  const changeIngredient = () => {
    setChangeIngredient({ change: true, id: entryData._id });
  };

  const changeEntry = () => {
    setChangeEntry({ change: true, id: entryData._id });
  };
  return (
    <Card>
      <Row>
        <Col>Dish:{entryData !== undefined && entryData.dish}</Col>
        <Col>
          Ingredients:{entryData !== undefined && entryData.ingredients}
        </Col>
        <Col>Calories:{entryData !== undefined && entryData.calories}</Col>
        <Col>Fat:{entryData !== undefined && entryData.fat}</Col>
        <Col>
          <Button
            onClick={() => {
              deleteEntry(entryData._id);
            }}
          >
            Delete
          </Button>
        </Col>
        <Col>
          <Button
            onClick={() => {
              changeIngredient();
            }}
          >
            Change Ingredients
          </Button>
        </Col>
        <Col>
          <Button
            onClick={() => {
              changeEntry();
            }}
          >
            Change Entry
          </Button>
        </Col>
      </Row>
    </Card>
  );
};

export default Entry;
