import React from 'react';
import {flexRender, getCoreRowModel, useReactTable} from '@tanstack/react-table';
import {columns} from "./columns";


const CoursesTable = ({ data }) => {
    const table = useReactTable({
        columns,
        data,
        getCoreRowModel: getCoreRowModel(),
    });

    return (
        <table>
            <thead >
            {table.getHeaderGroups().map(headerGroup => (
                <tr key={headerGroup.id}>
                    {headerGroup.headers.map(header => (
                        <th style={{width: '250px', border: '1px solid black'}} key={header.id}>
                            <h2>
                                {flexRender(header.column.columnDef.header, header.getContext())}
                            </h2>
                        </th>
                    ))}
                </tr>
            ))}
            </thead>
            <tbody>
            {table.getRowModel().rows.map(row => (
                <tr key={row.id}>
                    {row.getVisibleCells().map(cell => (
                        <td  key={cell.id}>
                            <h3 >
                                {flexRender(cell.column.columnDef.cell, cell.getContext())}
                            </h3>
                        </td>
                    ))}
                </tr>
            ))}
            </tbody>
        </table>
    );
};

export default CoursesTable;