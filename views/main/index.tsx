import React from "react";
import { HomeDataProps } from "../types/generate";

const Page = ({ title }: HomeDataProps) => {
  return (
    <div>
      <h4>{title}</h4>
    </div>
  );
};

export default Page;
