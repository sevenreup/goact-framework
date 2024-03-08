import React from "react";

export const Layout: React.FC = ({ children }) => {
  return (
    <div>
      <h1>My App</h1>
      {children}
    </div>
  );
};
