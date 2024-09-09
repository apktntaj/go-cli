"use client";

import React from "react";
import Alert from "./Alert";
import Row from "./Row";
import { isValidFormat } from "../utils/utility";
import TemplateExcelLink from "./TemplateExcelLink";

export default function Table({ data, setButtonStatus }) {
  console.log(data);
  if (!data) return <Alert message="Belum ada data" />;

  setButtonStatus("Disabled");

  return (
    <>
      <div>{}</div>
    </>
  );
}

function isFetched(data) {
  return data.every((item) => item["BM"]?.includes("%"));
}

function allHsCodesValid(data) {
  return data.every((item) => isValidFormat(item["HS CODE"]));
}
