import React, { useState } from "react";
import {
  flexRender,
  getCoreRowModel,
  useReactTable,
  getExpandedRowModel,
} from "@tanstack/react-table";
import { columns } from "./columns";
import styles from "../../shared/style/table.module.css";

const CoursesTable = ({ data, getRowCanExpand }) => {
  const table = useReactTable({
    columns,
    data,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <table className={styles["table"]}>
      <thead>
        {table.getHeaderGroups().map((headerGroup) => (
          <tr key={headerGroup.id}>
            {headerGroup.headers.map((header) => (
              <th
                // style={{ width: "250px", border: "1px solid #a7a7a7" }}
                key={header.id}
              >
                <h2>
                  {flexRender(
                    header.column.columnDef.header,
                    header.getContext()
                  )}
                </h2>
              </th>
            ))}
          </tr>
        ))}
      </thead>
      <tbody className={styles["table-body"]}>
        {table.getRowModel().rows.map((row) => (
          <tr className={styles["tr"]} key={row.id}>
            {row.getVisibleCells().map((cell) => (
              <td className={styles["td"]} key={cell.id}>
                <h3>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </h3>
              </td>
            ))}
            {/* {row.getIsExpanded() && (
              <tr>
                <td colSpan={row.getVisibleCells().length}>
                  <pre style={{ fontSize: "10px" }}>
                    <code>{JSON.stringify(row.original, null, 2)}</code>
                  </pre>
                </td>
              </tr>
            )} */}
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default CoursesTable;
