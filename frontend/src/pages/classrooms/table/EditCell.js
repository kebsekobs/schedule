import { Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { useDeleteClassroomMutation } from "../api/useDeleteClassroomMutation";
import { useState } from "react";
import EditModal from "../modals/editModal.js";
import styles from '../../shared/style/table.module.css';

export function EditCell(props) {
  const deleteClassroomMutation = useDeleteClassroomMutation();
  const id = props.props.original.id;
  const original = props.props.original;

  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  const toggleEditModal = () => {
    setIsEditModalOpen(!isEditModalOpen);
  };
  function deleteClassroom() {
    if (window.confirm("Вы уверены, что хотите удалить аудиторию?"))
    deleteClassroomMutation.mutateAsync({id});
  }

  return (
    <div style={{padding: '5px'}}>
      <button
        className={styles["btns-wrapper"]}
        onClick={() => toggleEditModal(id)}
      >
        <p>Изменить</p>
        <Pencil1Icon />
      </button>
      <button className={styles["btns-wrapper"]} onClick={deleteClassroom}>
        <p style={{ color: "var(--warning-color)" }}>Удалить</p>
        <TrashIcon />
      </button>
      {isEditModalOpen && (
        <EditModal
          toggleModal={toggleEditModal}
          isOpen={isEditModalOpen}
          original={original}
        />
      )}
    </div>
  );
}
