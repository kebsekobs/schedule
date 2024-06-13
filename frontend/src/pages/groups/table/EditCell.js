import { Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { useDeleteGroupMutation } from "../api/useDeleteGroupMutation";
import { useState } from "react";
import EditModal from "../modals/EditModal";
import styles from '../../shared/style/table.module.css';
import { DeleteModal } from "../modals/deleteModal";

export function EditCell(props) {
  const deleteGroupMutation = useDeleteGroupMutation();
  const id = props.props.original.id;
  const original = props.props.original;
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDeleteModalOpen , setisDeleteModalOpen] = useState(false);

  const toggleEditModal = () => {
    setIsEditModalOpen(!isEditModalOpen);
  };
  function deleteGroup() {
    setisDeleteModalOpen(!isDeleteModalOpen);
  }

  return (
    <div style={{padding: '5px'}}>
      <div
        className={styles["btns-wrapper"]}
        onClick={() => toggleEditModal(id)}
      >
        <p>Изменить</p>
        <Pencil1Icon />
      </div>
      <div className={styles["btns-wrapper"]} onClick={deleteGroup}>
        <p style={{ color: "var(--warning-color)" }}>Удалить</p>
        <TrashIcon />
      </div>
      {isEditModalOpen && (
        <EditModal
          toggleModal={toggleEditModal}
          isOpen={isEditModalOpen}
          original={original}
        />
      )}
      {isDeleteModalOpen && (
        <DeleteModal
        deleteGroup={deleteGroup}
        isOpen={isDeleteModalOpen}
        original={original}/>
      )}
    </div>
  );
}
