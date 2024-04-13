import {Pencil1Icon, TrashIcon} from "@radix-ui/react-icons";
import {useDeleteGroupMutation} from "../api/useDeleteGroupMutation";
import {useState} from "react";
import EditModal from "../modals/EditModal";


export function EditCell(props) {
    const deleteGroupMutation = useDeleteGroupMutation();
    const id = props.props.original.id;

    const [isEditModalOpen, setIsEditModalOpen] = useState(false);

    const toggleEditModal = () => {
        setIsEditModalOpen(!isEditModalOpen);
    };
    function deleteGroup() {
        deleteGroupMutation.mutateAsync(id);
    }

    return (
        <>
            <Pencil1Icon style={{ cursor: 'pointer' }} onClick={() => toggleEditModal(id)} />
            <TrashIcon style={{ cursor: 'pointer' }} onClick={deleteGroup} />
            {isEditModalOpen && <EditModal toggleModal={toggleEditModal} isOpen={isEditModalOpen} id={id} />
            }
        </>
    );
}
