import React from 'react';
export const CellHelper = {
    renderRelatedGroups: (row) => {
        const relatedGroupsArray = row.relatedGroupsId;
        return (
            <div style={{ textAlign: 'center'}}>
                {relatedGroupsArray.map((group, index) => (
                    <React.Fragment key={index}>
                        {group}{' '}
                    </React.Fragment>
                ))}
            </div>
        );
    }
}