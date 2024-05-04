import {createColumnHelper} from '@tanstack/react-table';
import {CellHelper} from "./CellHelpers";

const {accessor, group} = createColumnHelper();

export const columns = [
    group({
        id: '@id',
        header: 'id',
        columns: [
            accessor('id', {
                header: '',
                size: 400,
                cell: data =><div style={{ textAlign: 'center'}}>{data.getValue()}</div>,
            })
        ]
    }),
    group({
        id: '@name',
        header: 'name',
        columns: [
            accessor('name', {
                header: '',
                size: 400,
                cell: data =><div style={{ textAlign: 'center'}}>{data.getValue()}</div>,
            })
        ]
    }),
    group({
        id: '@teachers',
        header: 'teachers',
        columns: [
            accessor('teachers', {
                header: '',
                size: 400,
                cell: data =><div style={{ textAlign: 'center'}}>{data.getValue()}</div>,
            })
        ]
    }),
    group({
        id: '@hours',
        header: 'hours',
        columns: [
            accessor('hours', {
                header: '',
                size: 400,
                cell: data =><div style={{ textAlign: 'center'}}>{data.getValue()}</div>,
            })
        ]
    }),
    group({
        id: '@relatedGroupsId',
        header: 'relatedGroupsId',
        columns: [
            accessor('relatedGroupsId', {
                header: '',
                size: 400,
                cell: data => CellHelper.renderRelatedGroups(data.row.original),
            })
        ]
    }),

];

